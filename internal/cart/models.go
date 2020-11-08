package cart

import (
	"errors"
	"strings"

	"cinemo.com/shoping-cart/framework/web/validator"
	"cinemo.com/shoping-cart/internal/coupons"
	"cinemo.com/shoping-cart/internal/discounts"
	"cinemo.com/shoping-cart/internal/products"
)

type applyCouponOnCartRequest struct {
	CouponName string `json:"coupon"`
	CartID     int64  `json:"cart_id"`
}

func (req *applyCouponOnCartRequest) Validate() error {
	var errorArray []string
	validator.CheckRule(&errorArray, len(req.CouponName) > 0, "coupon is mandatory")
	validator.CheckRule(&errorArray, req.CartID > 0, "cart_id should be greater than 0")
	if len(errorArray) > 0 {
		return errors.New(strings.Join(errorArray, "; "))
	}
	return nil
}

type removeCouponFromCartRequest struct {
	CouponID int64 `json:"coupon_id"`
	CartID   int64 `json:"cart_id"`
}

func (req *removeCouponFromCartRequest) Validate() error {
	var errorArray []string
	validator.CheckRule(&errorArray, req.CouponID > 0, "coupon should be greater than 0")
	validator.CheckRule(&errorArray, req.CartID > 0, "cart_id should be greater than 0")
	if len(errorArray) > 0 {
		return errors.New(strings.Join(errorArray, "; "))
	}
	return nil
}

// addCartItemRequest htpp request for adding item in user cart
type addCartItemRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

func (req addCartItemRequest) Validate() error {
	var errorArray []string
	validator.CheckRule(&errorArray, req.ProductID > 0, "product_id is mandatory")
	validator.CheckRule(&errorArray, req.Quantity > -1, "quantity should be greater than -1")
	if len(errorArray) > 0 {
		return errors.New(strings.Join(errorArray, "; "))
	}
	return nil
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
