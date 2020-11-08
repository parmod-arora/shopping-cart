package coupons

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"cinemo.com/shoping-cart/internal/discounts"
	"cinemo.com/shoping-cart/internal/orm"
	"cinemo.com/shoping-cart/internal/products"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

const orangeDiscountName = "Coupon discount on oranges 30%"
const orangeProductName = "Oranges"

// Coupon struct represts the coupon type
type Coupon struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	ExpireAt   time.Time `json:"expire_at"`
	IsExpired  bool      `json:"expire"`
	ProductID  int64     `json:"product_id"`
	DiscountID int64     `json:"discount_id"`
}

func (s *couponService) CreateCoupon(ctx context.Context, now time.Time) (*Coupon, error) {
	discountID, err := s.discountService.FetchDiscountIDByName(ctx, orangeDiscountName)
	if err != nil {
		return nil, err
	}
	productID, err := s.productService.RetrieveProductIDByName(ctx, orangeProductName)
	if err != nil {
		return nil, err
	}
	return createCoupon(ctx, s.db, now, discountID, productID)
}

func (s *couponService) RetrieveCouponByName(ctx context.Context, couponName string) (*Coupon, error) {
	ormCoupon, err := orm.Coupons(
		qm.Where(orm.CouponColumns.Name+"=?", couponName),
	).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	if ormCoupon == nil {
		return &Coupon{}, nil
	}
	return TransformOrmToModel(ormCoupon), nil
}

func (s *couponService) RetrieveCouponByID(ctx context.Context, ID int64) (*Coupon, error) {
	ormCoupon, err := orm.Coupons(
		qm.Where(orm.CouponColumns.ID+"=?", ID),
	).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	if ormCoupon == nil {
		return &Coupon{}, nil
	}
	return TransformOrmToModel(ormCoupon), nil
}

func (s *couponService) RetrieveCouponProduct(ctx context.Context, couponName string, productID int64, timestamp time.Time) (*Coupon, error) {
	return retrieveCoupon(ctx, s.db, couponName, productID, timestamp.UTC())
}

func retrieveCoupon(ctx context.Context, db *sql.DB, couponName string, productID int64, timestamp time.Time) (*Coupon, error) {
	ormCoupon, err := orm.Coupons(
		qm.Where(orm.CouponColumns.Name+"=?", couponName),
		qm.And(orm.CouponColumns.ProductID+"=?", productID),
		qm.And(orm.CouponColumns.ExpireAt+" > ?", timestamp),
		qm.And(orm.CouponColumns.RedeemedAt+" is NULL"),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}
	if ormCoupon == nil {
		return &Coupon{}, nil
	}
	return TransformOrmToModel(ormCoupon), nil
}

func createCoupon(ctx context.Context, db *sql.DB, timestamp time.Time, discountID, productID int64) (*Coupon, error) {
	timestampStr := strconv.FormatInt(timestamp.UnixNano(), 10)
	ormCoupon := &orm.Coupon{
		Name:       "ORANGE_" + timestampStr,
		ProductID:  productID,
		DiscountID: discountID,
		ExpireAt:   timestamp.Add(time.Hour),
	}
	err := ormCoupon.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return TransformOrmToModel(ormCoupon), nil
}

// TransformOrmToModel transform orm to model
func TransformOrmToModel(coupon *orm.Coupon) *Coupon {
	return &Coupon{
		ID:         coupon.ID,
		Name:       coupon.Name,
		DiscountID: coupon.DiscountID,
		ProductID:  coupon.ProductID,
		IsExpired:  coupon.ExpireAt.Before(time.Now().UTC()),
		ExpireAt:   coupon.ExpireAt,
	}
}
