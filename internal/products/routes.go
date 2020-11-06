package products

import (
	"net/http"

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
	r.Path("/").Methods(http.MethodGet).HandlerFunc(ListProduct(service, userService))
}

// ListProduct ListProduct
func ListProduct(service Service, userService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// @TODO create a interceptor for this task
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

		// retrieve products from DB
		products, err := service.RetrieveProducts(ctx)
		if err != nil {
			status, errCode := statusAndErrorCodeForServiceError(err)
			httpresponse.ErrorResponseJSON(ctx, w, status, errCode, err.Error())
			return
		}

		httpresponse.RespondJSON(w, http.StatusOK, products, nil)
	}
}
