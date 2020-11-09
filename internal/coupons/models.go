package coupons

import (
	"time"

	"cinemo.com/shoping-cart/internal/orm"
)

// Coupon struct represts the coupon type
type Coupon struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	ExpireAt   time.Time  `json:"expire_at"`
	RedeemedAt *time.Time `json:"redeemed_at,omitempty"`
	IsExpired  bool       `json:"expire"`
	ProductID  int64      `json:"product_id"`
	DiscountID int64      `json:"discount_id"`
}

// TransformOrmToModel transform orm to model
func TransformOrmToModel(coupon *orm.Coupon) *Coupon {
	return &Coupon{
		ID:         coupon.ID,
		Name:       coupon.Name,
		DiscountID: coupon.DiscountID,
		ProductID:  coupon.ProductID,
		RedeemedAt: coupon.RedeemedAt.Ptr(),
		IsExpired:  coupon.ExpireAt.Before(time.Now().UTC()),
		ExpireAt:   coupon.ExpireAt,
	}
}
