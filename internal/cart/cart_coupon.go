package cart

import (
	"context"
	"errors"

	"cinemo.com/shoping-cart/internal/discounts"
	"cinemo.com/shoping-cart/internal/orm"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var errExpireCoupon error = errors.New("coupon_expired")

// GetCouponDiscountOnCart apply coupon on cart
func (s *cartService) getCouponDiscountOnCart(ctx context.Context, userCart *UserCart, couponID int64) (*discounts.ProductDiscount, error) {
	coupon, err := s.couponService.RetrieveCouponByID(ctx, couponID)
	if err != nil {
		return nil, err
	}

	// validate coupon expiry
	if coupon.IsExpired {
		return nil, nil
	}

	discount, err := s.discountService.FetchDiscountByID(ctx, coupon.DiscountID)
	if err != nil {
		return nil, err
	}

	return &discount, err
}

// RemoveCouponFromCart remove coupon from cart
func (s *cartService) RemoveCouponFromCart(ctx context.Context, couponID, cartID int64) error {
	_, err := orm.CartCoupons(
		qm.Where(orm.CartCouponColumns.CartID+"=?", cartID),
		qm.And(orm.CartCouponColumns.CouponID+"=?", couponID),
	).DeleteAll(ctx, s.db)
	return err
}
