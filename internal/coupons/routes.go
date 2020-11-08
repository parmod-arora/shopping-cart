package coupons

import (
	"net/http"
	"time"

	"cinemo.com/shoping-cart/framework/web/httpresponse"
	"cinemo.com/shoping-cart/framework/web/middleware"
	"cinemo.com/shoping-cart/internal/errorcode"
	"cinemo.com/shoping-cart/internal/users"
	"cinemo.com/shoping-cart/pkg/auth"
	"github.com/gorilla/mux"
)

// Handlers handles users routes
func Handlers(r *mux.Router, service Service, userService users.Service) {
	r.Use(middleware.Authorize)
	r.Path("/").Methods(http.MethodPut).HandlerFunc(CreateCoupon(service, userService))
}

// CreateCoupon creates new coupon in db
func CreateCoupon(service Service, userService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		username, err := auth.GetLoggedInUsername(r)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusForbidden, errorcode.AuthValidateFailed, err.Error())
			return
		}

		user, err := userService.RetrieveUserByUsername(ctx, username)
		if err != nil || user == nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusForbidden, errorcode.AuthValidateFailed, err.Error())
			return
		}

		coupon, err := service.CreateCoupon(ctx, time.Now().UTC())
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusInternalServerError, "internal_error", err.Error())
			return
		}
		httpresponse.RespondJSON(w, http.StatusCreated, coupon, nil)
	}
}
