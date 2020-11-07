package cart

import (
	"context"
	"database/sql"
	"math"

	"cinemo.com/shoping-cart/framework/loglib"
	"cinemo.com/shoping-cart/internal/discount"
	"cinemo.com/shoping-cart/internal/orm"
	"cinemo.com/shoping-cart/internal/products"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Service is the interface to expose order functions
type Service interface {
	GetUserCart(ctx context.Context, userID int64) (*UserCart, error)
	AddItemCart(ctx context.Context, userID int64, productID int64, quantity int64) (*UserCart, error)
}

type cartService struct {
	DB              *sql.DB
	DiscountService discount.Service
}

// NewCartService Create New cart service
func NewCartService(db *sql.DB, service discount.Service) Service {
	return cartService{
		DB:              db,
		DiscountService: service,
	}
}

// UserCart user cart for checkout page
type UserCart struct {
	ID                int64      `json:"id,omitempty"`
	UserID            int64      `json:"user_id,omitempty"`
	SubTotalAmount    int64      `json:"sub_total_amount,omitempty"`
	TotalSavingAmount int64      `json:"total_saving_amount,omitempty"`
	TotalAmount       int64      `json:"total_amount,omitempty"`
	CartItems         []CartItem `json:"cart_items,omitempty"`
	LineItems         []LineItem `json:"line_items,omitempty"`
}

// CartItem items added by user
type CartItem struct {
	ID       int64            `json:"id,omitempty"`
	CartID   int64            `json:"cart_id,omitempty"`
	Product  products.Product `json:"product,omitempty"`
	Quantity int64            `json:"quantity,omitempty"`
	SubTotal int64            `json:"sub_total"`
}

// AddItemCart add product in cart by user
func (s cartService) AddItemCart(ctx context.Context, userID int64, productID int64, quantity int64) (*UserCart, error) {
	logger := loglib.GetLogger(ctx)
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	// get or create user cart entry
	cartID, err := getORCreateCartEntry(ctx, tx, userID)
	if err != nil {
		logger.Errorf("Error %v", err.Error())
		return nil, err
	}
	if quantity == 0 {
		if err := deleteUserCartItemEntry(ctx, tx, cartID, productID); err != nil {
			logger.Errorf("Error deleteUserCartItemEntry %v", err.Error())
			return nil, err
		}
	} else {
		// add cart item entry
		if err := createUserCartItemEntry(ctx, tx, cartID, productID, quantity); err != nil {
			logger.Errorf("Error createUserCartItemEntry %v", err.Error())
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		logger.Errorf("Error commit%v", err.Error())
		return nil, err
	}
	return s.GetUserCart(ctx, userID)
}

func itemQuantityMatchesWithDiscountQuantity(discountQuantity discount.Quantity, itemQuantity int64) bool {
	switch discountQuantity.Function {
	case discount.EQ:
		if discountQuantity.Value == itemQuantity {
			return true
		}
	case discount.GTE:
		if discountQuantity.Value <= itemQuantity {
			return true
		}
	case discount.LTE:
		if discountQuantity.Value >= itemQuantity {
			return true
		}
	}
	return false
}

func (s cartService) GetUserCart(ctx context.Context, userID int64) (*UserCart, error) {
	cart, err := orm.Carts(
		qm.Load(qm.Rels(orm.CartRels.CartItems, orm.CartItemRels.Product)),
		qm.Where(orm.CartColumns.UserID+"=?", userID),
	).One(ctx, s.DB)
	if err != nil {
		if err == sql.ErrNoRows {
			return &UserCart{}, nil
		}
		return nil, err
	}
	usercart := &UserCart{
		ID:                cart.ID,
		UserID:            cart.UserID,
		TotalAmount:       int64(0),
		SubTotalAmount:    int64(0),
		TotalSavingAmount: int64(0),
	}
	subTotal := int64(0)
	cartItems := cart.R.CartItems
	for _, cartItem := range cartItems {
		subTotal = subTotal + (cartItem.R.Product.Amount * cartItem.Quantity)
		usercart.CartItems = append(usercart.CartItems, transformOrmToCartItem(cartItem))
	}
	usercart.SubTotalAmount = subTotal

	usercart, err = applyDiscountRules(ctx, usercart, s.DiscountService)
	if err != nil {
		return nil, err
	}

	// total saving
	for _, lineItem := range usercart.LineItems {
		usercart.TotalSavingAmount = usercart.TotalSavingAmount + lineItem.DiscountAmount
	}

	usercart.TotalAmount = usercart.SubTotalAmount - usercart.TotalSavingAmount
	return usercart, nil
}

func transformOrmToCartItem(cartItem *orm.CartItem) CartItem {
	return CartItem{
		ID: cartItem.ID,
		Product: products.Product{
			ID:        cartItem.R.Product.ID,
			Amount:    cartItem.R.Product.Amount,
			CreatedAt: cartItem.R.Product.CreatedAt,
			Details:   cartItem.R.Product.Details,
			Name:      cartItem.R.Product.Name,
			Image:     cartItem.R.Product.Image,
			UpdatedAt: cartItem.R.Product.UpdatedAt,
		},
		CartID:   cartItem.CartID,
		SubTotal: cartItem.Quantity * cartItem.R.Product.Amount,
		Quantity: cartItem.Quantity,
	}
}

func createUserCartItemEntry(ctx context.Context, db boil.ContextExecutor, cartID, productID, quantity int64) error {
	cart := &orm.CartItem{
		CartID:    cartID,
		ProductID: productID,
		Quantity:  quantity,
	}
	if err := cart.Upsert(ctx, db, true, []string{
		orm.CartItemColumns.CartID,
		orm.CartItemColumns.ProductID,
	}, boil.Whitelist(
		orm.CartItemColumns.Quantity,
		orm.CartItemColumns.UpdatedAt,
	), boil.Infer()); err != nil {
		return err
	}
	return nil
}

func deleteUserCartItemEntry(ctx context.Context, db boil.ContextExecutor, cartID, productID int64) error {
	_, err := orm.CartItems(qm.Where(orm.CartItemColumns.CartID+"=? AND "+orm.CartItemColumns.ProductID+"=?", cartID, productID)).DeleteAll(ctx, db)
	return err
}

func getORCreateCartEntry(ctx context.Context, db boil.ContextExecutor, userID int64) (int64, error) {
	cart := &orm.Cart{
		UserID:    userID,
		Reference: uuid.New().String(),
	}
	err := cart.Upsert(ctx, db, true, []string{
		orm.CartColumns.UserID,
	}, boil.Whitelist(
		orm.CartColumns.UpdatedAt,
	), boil.Infer())
	if err != nil {
		return 0, err
	}
	return cart.ID, nil
}

// LineItem lineitem for checkout page
type LineItem struct {
	ProductDiscount *discount.ProductDiscount `json:"discount_applied"`
	CartItems       []CartItem                `json:"discount_applied_on"`
	DiscountAmount  int64                     `json:"discount_amount"`
	Quantity        int64                     `json:"quantity"`
}

// ComboCartItem combo cart item
type ComboCartItem struct {
	CartItem           CartItem `json:"cart_item,omitempty"`
	PackedWithCartItem CartItem `json:"packed_with_cart_item,omitempty"`
}

func applyDiscountRules(ctx context.Context, usercart *UserCart, discountService discount.Service) (*UserCart, error) {
	logger := loglib.GetLogger(ctx)
	productDiscountMap := make(map[int64][]discount.ProductDiscount)
	productMap := make(map[int64]*CartItem)
	for _, item := range usercart.CartItems {
		item := item
		if _, ok := productMap[item.Product.ID]; !ok {
			productMap[item.Product.ID] = &item
			discounts, err := discountService.RetrieveProductDiscounts(ctx, item.Product.ID)
			if err != nil {
				return nil, err
			}
			productDiscountMap[item.Product.ID] = discounts
		}
	}

	logger.Infof("productDiscountMap %v", productDiscountMap)

	// quantity based product discount
	for _, item := range usercart.CartItems {
		var discountFound *discount.ProductDiscount
		productDiscounts, ok := productDiscountMap[item.Product.ID]
		if ok {
			// find first product discount and break loop
			for _, productDiscount := range productDiscounts {
				if len(productDiscount.Rules) == 1 {
					rule := productDiscount.Rules[0]
					ruleQuantity := discount.Quantity{
						Function: rule.ProductQuantityFN,
						Value:    rule.ProductQuantity,
					}
					if itemQuantityMatchesWithDiscountQuantity(ruleQuantity, item.Quantity) {
						discountFound = &productDiscount
						break
					}
				}
			}
			if discountFound != nil {
				// discountFound.Discount
				lineItem := LineItem{
					ProductDiscount: discountFound,
				}
				lineItem.CartItems = append(lineItem.CartItems, item)
				if discountFound.DiscountType == discount.PERCENTAGE {
					lineItem.DiscountAmount = int64(math.Round(float64(discountFound.Discount)/float64(100)*float64(item.Product.Amount))) * item.Quantity
				}
				lineItem.Quantity = 1
				usercart.LineItems = append(usercart.LineItems, lineItem)
				// remove items from map if discount is applied on those
				delete(productMap, item.Product.ID)
			}
		}
	}

	// iterate remaining product for combo discount
	for productID := range productMap {
		if productDiscounts, ok := productDiscountMap[productID]; ok {
			var discountFound *discount.ProductDiscount
			for _, productDiscount := range productDiscounts {
				// does cart satisfy discount rule
				numberOfRulesMatch := 0
				for _, rule := range productDiscount.Rules {
					//  cart contains discount product
					cartItem, ok := productMap[rule.ProductID]
					if !ok {
						break
					}
					ruleQuantity := discount.Quantity{
						Function: discount.GTE,
						Value:    rule.ProductQuantity,
					}
					if !itemQuantityMatchesWithDiscountQuantity(ruleQuantity, cartItem.Quantity) {
						break
					}
					numberOfRulesMatch = numberOfRulesMatch + 1
				}
				if numberOfRulesMatch == len(productDiscount.Rules) {
					discountFound = &productDiscount
				}
			}

			if discountFound != nil {
				lineItem := LineItem{
					ProductDiscount: discountFound,
				}
				// discountFound.Discount
				if discountFound.DiscountType == discount.PERCENTAGE {
					noOfSet := math.MaxFloat64
					subTotal := int64(0)
					for _, rule := range discountFound.Rules {
						cartItem := productMap[rule.ProductID]
						subTotal = subTotal + (cartItem.Product.Amount * rule.ProductQuantity)
						noOfSet = math.Min(noOfSet, float64(cartItem.Quantity/rule.ProductQuantity))
						lineItem.CartItems = append(lineItem.CartItems, *cartItem)
					}
					lineItem.Quantity = int64(noOfSet)
					for _, rule := range discountFound.Rules {
						cartItem := productMap[rule.ProductID]
						cartItem.Quantity = cartItem.Quantity - (rule.ProductQuantity * lineItem.Quantity)
					}
					lineItem.DiscountAmount = int64(math.Round(float64(discountFound.Discount) / float64(100) * float64(subTotal) * noOfSet))
					usercart.LineItems = append(usercart.LineItems, lineItem)
				}
			}
		}
	}

	return usercart, nil
}
