package coupons

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"cinemo.com/shoping-cart/internal/errorcode"
	"cinemo.com/shoping-cart/internal/orm"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const orangeDiscountName = "30% coupon discount on oranges"
const orangeProductName = "Oranges"

func (s *couponService) CreateCoupon(ctx context.Context, now time.Time) (*Coupon, error) {
	discountID, err := s.discountService.FetchDiscountIDByName(ctx, orangeDiscountName)
	if err != nil {
		return nil, errorcode.DBError{Err: err}
	}
	productID, err := s.productService.RetrieveProductIDByName(ctx, orangeProductName)
	if err != nil {
		return nil, errorcode.DBError{Err: err}
	}
	return createCoupon(ctx, s.db, now, discountID, productID)
}

func (s *couponService) RetrieveCouponByName(ctx context.Context, couponName string) (*Coupon, error) {
	ormCoupon, err := orm.Coupons(
		qm.Where(orm.CouponColumns.Name+"=?", couponName),
	).One(ctx, s.db)
	if err != nil {
		return nil, errorcode.DBError{Err: err}
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
		return nil, errorcode.DBError{Err: err}
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
		return nil, errorcode.DBError{Err: err}
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
		return nil, errorcode.DBError{Err: err}
	}
	return TransformOrmToModel(ormCoupon), nil
}
