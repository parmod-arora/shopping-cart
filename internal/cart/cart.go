package cart

import (
	"context"
	"database/sql"
	"encoding/json"
	"math"
	"time"

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

type UserCart struct {
	ID                int64
	UserID            int64
	TotalAmount       int64
	TotalSavingAmount int64
	CartItems         []CartItem
	LineItems         []LineItem
}

type CartItem struct {
	ID        int64
	CartID    int64
	Product   products.Product
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AddItemCart add product in cart by user
func (s cartService) AddItemCart(ctx context.Context, userID int64, productID int64, quantity int64) (*UserCart, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	// get or create user cart entry
	cartID, err := getORCreateCartEntry(ctx, tx, userID)
	if err != nil {
		return nil, err
	}
	// add cart item entry

	if err := createUserCartItemEntry(ctx, tx, userID, cartID, quantity); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
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
		return nil, err
	}
	usercart := &UserCart{
		ID:     cart.ID,
		UserID: cart.UserID,
	}
	totalAmount := int64(0)
	cartItems := cart.R.CartItems
	for _, cartItem := range cartItems {
		totalAmount = totalAmount + (cartItem.R.Product.Amount * cartItem.Quantity)
		usercart.CartItems = append(usercart.CartItems, CartItem{
			ID: cartItem.ID,
			Product: products.Product{
				ID:        cartItem.R.Product.ID,
				Amount:    cartItem.R.Product.Amount,
				CreatedAt: cartItem.R.Product.CreatedAt,
				Details:   cartItem.R.Product.Details,
				Name:      cartItem.R.Product.Name,
				UpdatedAt: cartItem.R.Product.UpdatedAt,
			},
			Quantity:  cartItem.Quantity,
			CreatedAt: cartItem.CreatedAt,
			UpdatedAt: cartItem.UpdatedAt,
		})
	}

	usercart, err = applyDiscountRules(ctx, usercart, s.DiscountService)
	if err != nil {
		return nil, err
	}

	usercart.TotalAmount = totalAmount
	return usercart, nil
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

func getORCreateCartEntry(ctx context.Context, db boil.ContextExecutor, userID int64) (int64, error) {
	cart := &orm.Cart{
		UserID:    userID,
		Reference: uuid.New().String(),
	}
	err := cart.Upsert(ctx, db, true, []string{}, boil.Whitelist(
		orm.CartColumns.UpdatedAt,
	), boil.Infer())
	if err != nil {
		return 0, err
	}
	return cart.ID, nil
}

// LineItem lineitem for checkout page
type LineItem struct {
	CartItem         CartItem
	ComboCartItem    ComboCartItem
	ComboDiscount    *discount.ComboDiscountProduct
	Discount         *discount.ProductDiscount
	Quantity         int64
	Amount           int64
	DiscountedAmount int64
	SubTotal         int64
}

type ComboCartItem struct {
	CartItem           CartItem
	PackedWithCartItem CartItem
}

func applyDiscountRules(ctx context.Context, usercart *UserCart, discountService discount.Service) (*UserCart, error) {
	logger := loglib.GetLogger(ctx)
	productDiscountMap := make(map[int64]map[int64]discount.Rules)
	productMap := make(map[int64]*CartItem)
	for _, item := range usercart.CartItems {
		item := item
		if _, ok := productDiscountMap[item.Product.ID]; !ok {
			productMap[item.Product.ID] = &item
			rules, err := discountService.RetrieveProductDiscounts(ctx, item.Product.ID)
			if err != nil {
				return nil, err
			}
			productDiscountMap[item.Product.ID] = rules
		}
	}

	bytes, _ := json.Marshal(productDiscountMap)
	logger.Infof("%v", string(bytes))

	// quantity based product discount
	for _, item := range usercart.CartItems {
		var discountFound *discount.ProductDiscount
		productDiscountRules, _ := productDiscountMap[item.Product.ID]
		rules, _ := productDiscountRules[item.Product.ID]
		productDiscounts := rules.ProductDiscounts

		lineitem := LineItem{
			CartItem:         item,
			DiscountedAmount: item.Product.Amount,
			Quantity:         1,
			SubTotal:         item.Quantity * item.Product.Amount,
		}

		// find first product discount and break loop
		for _, productDiscount := range productDiscounts {
			if itemQuantityMatchesWithDiscountQuantity(productDiscount.Quantity, item.Quantity) {
				discountFound = &productDiscount
				break
			}
		}

		if discountFound != nil {
			lineitem.Discount = discountFound
			// apply percentage discount rule
			if discountFound.DiscountType == discount.PERCENTAGE {
				discountedAmount := float64(item.Product.Amount) - float64(discountFound.Discount)/float64(100)*float64(item.Product.Amount)
				lineitem.DiscountedAmount = int64(math.Round(discountedAmount))
				lineitem.Amount = item.Quantity * item.Product.Amount
				lineitem.SubTotal = lineitem.DiscountedAmount * item.Quantity
			}
			usercart.LineItems = append(usercart.LineItems, lineitem)
			// remove items from map if discount is applied on those
			delete(productMap, item.Product.ID)
		}
	}

	// iterate remaining product for combo discount
	for productID, cartItem := range productMap {
		if rulesMap, ok := productDiscountMap[productID]; ok {
			if rules, hasRules := rulesMap[productID]; hasRules {
				// this product has combo discount rule
				// check cartitems against combo rule
				for packagedWithProdcutID, comboDiscount := range rules.ComboDiscounts {
					packagedWithCartItem := productMap[packagedWithProdcutID]
					// check combo discount conditions
					// check conditions for product and packaged with product
					if itemQuantityMatchesWithDiscountQuantity(discount.Quantity{Function: discount.GTE, Value: comboDiscount.Quantity.Value}, cartItem.Quantity) &&
						itemQuantityMatchesWithDiscountQuantity(discount.Quantity{Function: discount.GTE, Value: comboDiscount.PackagedWithQuantity.Value}, packagedWithCartItem.Quantity) {

						// prepare combo line item
						lineItemCount := int64(0)
						for cartItem.Quantity >= comboDiscount.Quantity.Value && packagedWithCartItem.Quantity >= comboDiscount.PackagedWithQuantity.Value {
							cartItem.Quantity = cartItem.Quantity - comboDiscount.Quantity.Value
							packagedWithCartItem.Quantity = packagedWithCartItem.Quantity - comboDiscount.PackagedWithQuantity.Value
							lineItemCount = lineItemCount + 1
						}
						lineitem := LineItem{
							ComboCartItem: ComboCartItem{
								CartItem: CartItem{
									CartID:   cartItem.CartID,
									Product:  cartItem.Product,
									Quantity: comboDiscount.Quantity.Value,
								},
								PackedWithCartItem: CartItem{
									CartID:   packagedWithCartItem.CartID,
									Product:  packagedWithCartItem.Product,
									Quantity: comboDiscount.PackagedWithQuantity.Value,
								},
							},
							ComboDiscount: &comboDiscount,
							Quantity:      lineItemCount,
						}
						if comboDiscount.DiscountType == discount.PERCENTAGE {
							productAmount := int64(math.Round(float64(cartItem.Product.Amount) - float64(comboDiscount.Discount)/float64(100)*float64(cartItem.Product.Amount)))
							packagedWithProductAmount := int64(math.Round(float64(packagedWithCartItem.Product.Amount) - float64(comboDiscount.Discount)/float64(100)*float64(packagedWithCartItem.Product.Amount)))
							lineitem.SubTotal = productAmount*comboDiscount.Quantity.Value + packagedWithProductAmount*comboDiscount.PackagedWithQuantity.Value
							lineitem.Amount = cartItem.Product.Amount*comboDiscount.Quantity.Value + packagedWithCartItem.Product.Amount*comboDiscount.PackagedWithQuantity.Value
						}
						usercart.LineItems = append(usercart.LineItems, lineitem)
					}
				}
			}
		}
	}

	// prepare line entry for remaining items
	for _, cartItem := range productMap {
		if cartItem.Quantity > 0 {
			lineitem := LineItem{
				CartItem: *cartItem,
				Amount:   cartItem.Product.Amount,
				SubTotal: cartItem.Product.Amount * cartItem.Quantity,
			}
			usercart.LineItems = append(usercart.LineItems, lineitem)
		}
	}

	return usercart, nil
}