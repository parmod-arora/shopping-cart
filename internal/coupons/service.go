package coupons

import (
	"context"
	"database/sql"
	"time"

	"cinemo.com/shoping-cart/internal/discounts"
	"cinemo.com/shoping-cart/internal/products"
)

// Service is the interface to expose coupons functions
type Service interface {
	CreateCoupon(ctx context.Context, now time.Time) (*Coupon, error)
	RetrieveCouponProduct(ctx context.Context, couponName string, productID int64, timestamp time.Time) (*Coupon, error)
	RetrieveCouponByName(ctx context.Context, couponName string) (*Coupon, error)
	RetrieveCouponByID(ctx context.Context, ID int64) (*Coupon, error)
}

type couponService struct {
	db              *sql.DB
	productService  products.Service
	discountService discounts.Service
}

// NewCouponService creates new coupon service
func NewCouponService(db *sql.DB, productService products.Service, discountService discounts.Service) Service {
	return &couponService{
		db:              db,
		productService:  productService,
		discountService: discountService,
	}
}
