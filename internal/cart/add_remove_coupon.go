package cart

import (
	"encoding/json"
	"net/http"

	"cinemo.com/shoping-cart/framework/web/httpresponse"
	"cinemo.com/shoping-cart/internal/errorcode"
	"cinemo.com/shoping-cart/internal/users"
	"cinemo.com/shoping-cart/pkg/auth"
)

// ApplyCouponOnCart apply coupon on cart
func ApplyCouponOnCart(service Service, userService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		username, err := auth.GetLoggedInUsername(r)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusForbidden, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		user, err := userService.RetrieveUserByUsername(ctx, username)
		if err != nil || user == nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		// unmarshal request
		req := applyCouponOnCartRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); (err != nil || req == applyCouponOnCartRequest{}) {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		// validate request
		if err := req.Validate(); err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		if err := service.ApplyCouponOnCart(ctx, req.CouponName, req.CartID, user.ID); err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusInternalServerError, errorcode.InternalError, err.Error())
			return
		}

		cart, err := service.GetUserCart(ctx, user.ID)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusInternalServerError, "internal_error", err.Error())
			return
		}
		httpresponse.RespondJSON(w, http.StatusOK, cart, nil)
	}
}

// RemoveCouponFromCart remove coupon from cart
func RemoveCouponFromCart(service Service, userService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		username, err := auth.GetLoggedInUsername(r)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusForbidden, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		user, err := userService.RetrieveUserByUsername(ctx, username)
		if err != nil || user == nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		// unmarshal request
		req := removeCouponFromCartRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); (err != nil || req == removeCouponFromCartRequest{}) {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		// validate request
		if err := req.Validate(); err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		if err := service.RemoveCouponFromCart(ctx, req.CouponID, req.CartID); err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusInternalServerError, errorcode.InternalError, err.Error())
			return
		}

		cart, err := service.GetUserCart(ctx, user.ID)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusInternalServerError, "internal_error", err.Error())
			return
		}
		httpresponse.RespondJSON(w, http.StatusOK, cart, nil)
	}
}
