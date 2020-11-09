package cart

import (
	"context"
	"database/sql"

	"cinemo.com/shoping-cart/internal/coupons"
	"cinemo.com/shoping-cart/internal/discounts"
)

// Service is the interface to expose order functions
type Service interface {
	GetUserCart(ctx context.Context, userID int64) (*UserCart, error)
	AddItemCart(ctx context.Context, userID int64, productID int64, quantity int64) (*UserCart, error)
	ApplyCouponOnCart(ctx context.Context, couponName string, cartID, userID int64) error
	RemoveCouponFromCart(ctx context.Context, couponID, cartID int64) error
	CheckoutCart(ctx context.Context, usercart *UserCart) error
}

type cartService struct {
	db              *sql.DB
	discountService discounts.Service
	couponService   coupons.Service
}

// NewCartService Create New cart service
func NewCartService(db *sql.DB, service discounts.Service, couponService coupons.Service) Service {
	return &cartService{
		db:              db,
		discountService: service,
		couponService:   couponService,
	}
}
