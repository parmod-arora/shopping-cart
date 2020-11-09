package cart

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"cinemo.com/shoping-cart/framework/web/httpresponse"
	"cinemo.com/shoping-cart/internal/errorcode"
	"cinemo.com/shoping-cart/internal/orm"
	"cinemo.com/shoping-cart/internal/users"
	"cinemo.com/shoping-cart/pkg/auth"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type checkoutRequest struct {
	CartID int64 `json:"cart_id"`
	Amount int64 `json:"amount"`
}

type checkoutResponse struct {
	OrderReference int64 `json:"order_reference"`
	Amount         int64 `json:"amount"`
}

// Checkout retrieve User cart from DB
func Checkout(service Service, userService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username, err := auth.GetLoggedInUsername(r)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusForbidden, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		user, err := userService.RetrieveUserByUsername(ctx, username)
		if err != nil || user == nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusUnauthorized, errorcode.UserNotFound, "User not found")
			return
		}

		// unmarshal request
		req := checkoutRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); (err != nil || req == checkoutRequest{}) {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		// validate request
		if req.CartID <= 0 {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		usercart, err := service.GetUserCart(ctx, user.ID)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusInternalServerError, errorcode.InternalError, err.Error())
			return
		}

		if usercart.TotalAmount != req.Amount {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusConflict, errorcode.AmountMisMatch, "Amount mismatch")
			return
		}

		err = service.CheckoutCart(ctx, usercart)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusInternalServerError, errorcode.InternalError, err.Error())
			return
		}

		httpresponse.RespondJSON(w, http.StatusOK, nil, nil)
	}
}

// CheckoutCart create order and empty user cart
func (s *cartService) CheckoutCart(ctx context.Context, usercart *UserCart) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := transferCart(ctx, tx, usercart); err != nil {
		return err
	}
	return tx.Commit()
}

// transferCart create order from cart enteries
func transferCart(ctx context.Context, tx boil.ContextExecutor, usercart *UserCart) error {
	// create order from cart
	ormOrder := orm.Order{
		Reference:      uuid.New().String(),
		UserID:         usercart.UserID,
		TotalAmount:    usercart.TotalAmount,
		SubTotalAmount: usercart.SubTotalAmount,
	}
	if err := ormOrder.Insert(ctx, tx, boil.Infer()); err != nil {
		return err
	}

	for _, cartItem := range usercart.CartItems {
		ormOrderItem := orm.OrderItem{
			OrderID:   ormOrder.ID,
			ProductID: cartItem.Product.ID,
			Quantity:  cartItem.Quantity,
		}
		ormOrderItem.Insert(ctx, tx, boil.Infer())
	}

	// delete cart items
	if _, err := orm.CartItems(qm.Where(orm.CartItemColumns.CartID+"=?", usercart.ID)).DeleteAll(ctx, tx); err != nil {
		return err
	}

	if usercart.Coupon != nil {
		ormOrderCoupon := &orm.OrderCoupon{
			CouponID: usercart.Coupon.ID,
			OrderID:  ormOrder.ID,
		}
		if err := ormOrderCoupon.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}

		// delete cart coupons
		if _, err := orm.CartCoupons(qm.Where(orm.CartCouponColumns.CartID+"=?", usercart.ID)).DeleteAll(ctx, tx); err != nil {
			return err
		}

		// mark coupon as redemmed
		ormCoupon, err := orm.Coupons(qm.Where(orm.CouponColumns.ID+"=?", usercart.Coupon.ID)).One(ctx, tx)
		if err != nil {
			return err
		}
		ormCoupon.RedeemedAt = null.TimeFrom(time.Now().UTC())
		_, err = ormCoupon.Update(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	for _, lineItem := range usercart.LineItems {
		ormDiscount := orm.OrderDiscount{
			DiscountID: lineItem.ProductDiscount.ID,
			OrderID:    ormOrder.ID,
		}
		if err := ormDiscount.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	}

	// delete user cart
	if _, err := orm.Carts(qm.Where(orm.CartColumns.ID+"=?", usercart.ID)).DeleteAll(ctx, tx); err != nil {
		return err
	}

	return nil
}
