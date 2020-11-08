package cart

import (
	"context"
	"database/sql"
	"errors"
	"math"

	"cinemo.com/shoping-cart/framework/loglib"
	"cinemo.com/shoping-cart/internal/coupons"
	"cinemo.com/shoping-cart/internal/discounts"
	"cinemo.com/shoping-cart/internal/errorcode"
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
	ApplyCouponOnCart(ctx context.Context, couponName string, cartID, userID int64) error
	RemoveCouponFromCart(ctx context.Context, couponID, cartID int64) error
}

type cartService struct {
	DB              *sql.DB
	discountService discounts.Service
	couponService   coupons.Service
}

// NewCartService Create New cart service
func NewCartService(db *sql.DB, service discounts.Service, couponService coupons.Service) Service {
	return &cartService{
		DB:              db,
		discountService: service,
		couponService:   couponService,
	}
}

// UserCart user cart for checkout page
type UserCart struct {
	ID                int64           `json:"id,omitempty"`
	UserID            int64           `json:"user_id,omitempty"`
	SubTotalAmount    int64           `json:"sub_total_amount,omitempty"`
	TotalSavingAmount int64           `json:"total_saving_amount,omitempty"`
	TotalAmount       int64           `json:"total_amount,omitempty"`
	Coupon            *coupons.Coupon `json:"coupon,omitempty"`
	CartItems         []CartItem      `json:"cart_items,omitempty"`
	LineItems         []LineItem      `json:"line_items,omitempty"`
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
func (s *cartService) AddItemCart(ctx context.Context, userID int64, productID int64, quantity int64) (*UserCart, error) {
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

func itemQuantityMatchesWithDiscountQuantity(discountQuantity discounts.Quantity, itemQuantity int64) bool {
	switch discountQuantity.Function {
	case discounts.EQ:
		if discountQuantity.Value == itemQuantity {
			return true
		}
	case discounts.GTE:
		if discountQuantity.Value <= itemQuantity {
			return true
		}
	case discounts.LTE:
		if discountQuantity.Value >= itemQuantity {
			return true
		}
	}
	return false
}

func (s *cartService) GetUserCart(ctx context.Context, userID int64) (*UserCart, error) {
	boil.DebugMode = true
	cart, err := orm.Carts(
		qm.Load(qm.Rels(orm.CartRels.CartItems, orm.CartItemRels.Product)),
		qm.Load(qm.Rels(orm.CartRels.CartCoupons, orm.CartCouponRels.Coupon)),
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

	// for now only one - one to mapping with cart and coupon
	for _, cartCoupon := range cart.R.CartCoupons {
		usercart.Coupon = coupons.TransformOrmToModel(cartCoupon.R.Coupon)
	}

	usercart, err = s.applyDiscountsOnCart(ctx, usercart)
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
	ProductDiscount *discounts.ProductDiscount `json:"discount_applied"`
	CartItems       []CartItem                 `json:"discount_applied_on"`
	DiscountAmount  int64                      `json:"discount_amount"`
	Quantity        int64                      `json:"quantity"`
}

// ComboCartItem combo cart item
type ComboCartItem struct {
	CartItem           CartItem `json:"cart_item,omitempty"`
	PackedWithCartItem CartItem `json:"packed_with_cart_item,omitempty"`
}

func (s *cartService) applyDiscountsOnCart(ctx context.Context, usercart *UserCart) (*UserCart, error) {
	logger := loglib.GetLogger(ctx)
	productDiscountMap := make(map[int64][]discounts.ProductDiscount)
	productMap := make(map[int64]*CartItem)
	for _, item := range usercart.CartItems {
		item := item
		if _, ok := productMap[item.Product.ID]; !ok {
			productMap[item.Product.ID] = &item
			discounts, err := s.discountService.RetrieveProductDiscounts(ctx, item.Product.ID)
			if err != nil {
				return nil, err
			}
			productDiscountMap[item.Product.ID] = discounts
		}
	}

	if usercart.Coupon != nil && usercart.Coupon.ID > 0 {
		// get coupn discount on cart
		couponDiscount, err := s.getCouponDiscountOnCart(ctx, usercart, usercart.Coupon.ID)
		if err != nil {
			return nil, err
		}
		if couponDiscount != nil {
			for _, rule := range couponDiscount.Rules {
				discounts, ok := productDiscountMap[rule.ProductID]
				if ok {
					discounts = append(discounts, *couponDiscount)
				}
				productDiscountMap[rule.ProductID] = discounts
			}
		}
	}

	logger.Infof("productDiscountMap %#v", productDiscountMap)
	for productID, productDiscounts := range productDiscountMap {
		logger.Infof("%v, %v", productID, productDiscounts)
		for _, productDiscount := range productDiscounts {
			productDiscount := productDiscount
			applyDiscountOnCart(ctx, productMap, usercart, productDiscount)
		}
	}
	return usercart, nil
}

func applyDiscountOnCart(ctx context.Context, productMap map[int64]*CartItem, usercart *UserCart, discount discounts.ProductDiscount) (*UserCart, error) {
	logger := loglib.GetLogger(ctx)

	// apply product discount
	applySingleDiscount(ctx, productMap, usercart, discount)

	logger.Infof("%#v", productMap)
	// apply combo product discount
	applyComboDiscount(ctx, productMap, usercart, discount)

	return usercart, nil
}

func applySingleDiscount(ctx context.Context, productMap map[int64]*CartItem, usercart *UserCart, productDiscount discounts.ProductDiscount) *UserCart {
	if len(productDiscount.Rules) != 1 {
		return usercart
	}
	var validDiscount *discounts.ProductDiscount
	for _, item := range usercart.CartItems {
		// check discount is of single type
		rule := productDiscount.Rules[0]
		if rule.ProductID != item.Product.ID {
			continue
		}
		ruleQuantity := discounts.Quantity{
			Function: rule.ProductQuantityFN,
			Value:    rule.ProductQuantity,
		}
		if itemQuantityMatchesWithDiscountQuantity(ruleQuantity, item.Quantity) {
			validDiscount = &productDiscount
		}
		if validDiscount != nil {
			lineItem := LineItem{
				ProductDiscount: validDiscount,
			}
			lineItem.CartItems = append(lineItem.CartItems, item)
			if validDiscount.DiscountType == discounts.PERCENTAGE {
				lineItem.DiscountAmount = int64(math.Round(float64(validDiscount.Discount)/float64(100)*float64(item.Product.Amount))) * item.Quantity
			}
			lineItem.Quantity = 1
			usercart.LineItems = append(usercart.LineItems, lineItem)
			// remove items from map if discount is applied on those
			delete(productMap, item.Product.ID)
		}
	}
	return usercart
}

func applyComboDiscount(ctx context.Context, productMap map[int64]*CartItem, usercart *UserCart, productDiscount discounts.ProductDiscount) *UserCart {
	if len(productDiscount.Rules) <= 1 {
		return usercart
	}
	var validDiscount *discounts.ProductDiscount
	numberOfRulesMatch := int(0)
	for _, rule := range productDiscount.Rules {
		//  cart contains discount product
		cartItem, ok := productMap[rule.ProductID]
		if !ok {
			break
		}
		if rule.ProductID != cartItem.Product.ID {
			continue
		}
		ruleQuantity := discounts.Quantity{
			Function: discounts.GTE,
			Value:    rule.ProductQuantity,
		}
		if !itemQuantityMatchesWithDiscountQuantity(ruleQuantity, cartItem.Quantity) {
			break
		}
		numberOfRulesMatch = numberOfRulesMatch + 1
	}
	if numberOfRulesMatch == len(productDiscount.Rules) {
		validDiscount = &productDiscount
	}
	if validDiscount != nil {
		lineItem := LineItem{
			ProductDiscount: validDiscount,
		}
		// validDiscount.Discount
		if validDiscount.DiscountType == discounts.PERCENTAGE {
			noOfSet := math.MaxFloat64
			subTotal := int64(0)
			for _, rule := range validDiscount.Rules {
				cartItem := productMap[rule.ProductID]
				subTotal = subTotal + (cartItem.Product.Amount * rule.ProductQuantity)
				noOfSet = math.Min(noOfSet, float64(cartItem.Quantity/rule.ProductQuantity))
				lineItem.CartItems = append(lineItem.CartItems, *cartItem)
			}
			lineItem.Quantity = int64(noOfSet)
			for _, rule := range validDiscount.Rules {
				cartItem := productMap[rule.ProductID]
				cartItem.Quantity = cartItem.Quantity - ((rule.ProductQuantity * lineItem.Quantity) * lineItem.Quantity)
			}
			lineItem.DiscountAmount = int64(math.Round(float64(validDiscount.Discount) / float64(100) * float64(subTotal) * noOfSet))
			usercart.LineItems = append(usercart.LineItems, lineItem)
		}
	}
	return usercart
}

func (s *cartService) ApplyCouponOnCart(ctx context.Context, couponName string, cartID, userID int64) error {

	userCartID, err := getORCreateCartEntry(ctx, s.DB, userID)
	if err != nil {
		return err
	}

	if cartID != userCartID {
		return errorcode.ValidationError{Err: errors.New("invalid-cart-id")}
	}

	coupon, err := s.couponService.RetrieveCouponByName(ctx, couponName)
	if err != nil || coupon == nil {
		return errorcode.ValidationError{Err: errors.New("invalid-coupon")}
	}

	if coupon.IsExpired {
		return errorcode.ValidationError{Err: errors.New("expired-coupon")}
	}

	usercart, err := s.GetUserCart(ctx, userID)
	if err != nil {
		return err
	}

	productDiscount, err := s.discountService.FetchDiscountByID(ctx, coupon.DiscountID)
	if err != nil {
		return err
	}
	if !(len(productDiscount.Rules) > 0) {
		return errorcode.ValidationError{Err: errors.New("invalid-coupon-rules")}
	}

	if usercart == nil {
		return errorcode.ValidationError{Err: errors.New("coupon-not-aplicable")}
	}
	var validDiscount *discounts.ProductDiscount
	for _, item := range usercart.CartItems {
		// check discount is of single type
		rule := productDiscount.Rules[0]
		ruleQuantity := discounts.Quantity{
			Function: rule.ProductQuantityFN,
			Value:    rule.ProductQuantity,
		}
		if itemQuantityMatchesWithDiscountQuantity(ruleQuantity, item.Quantity) {
			validDiscount = &productDiscount
			break
		}
	}

	if validDiscount == nil {
		return errorcode.ValidationError{Err: errors.New("coupon-not-aplicable")}
	}

	cartCoupon := &orm.CartCoupon{
		CartID:   cartID,
		CouponID: coupon.ID,
	}
	if err := cartCoupon.Insert(ctx, s.DB, boil.Infer()); err != nil {
		return err
	}
	return nil
}
