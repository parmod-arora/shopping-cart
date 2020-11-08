package cart

import (
	"context"
	"errors"

	"cinemo.com/shoping-cart/internal/discounts"
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
