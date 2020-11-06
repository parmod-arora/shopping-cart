package cart

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"cinemo.com/shoping-cart/framework/loglib"
	"cinemo.com/shoping-cart/framework/web/httpresponse"
	"cinemo.com/shoping-cart/framework/web/middleware"
	"cinemo.com/shoping-cart/framework/web/validator"
	"cinemo.com/shoping-cart/internal/errorcode"
	"cinemo.com/shoping-cart/internal/users"
	"cinemo.com/shoping-cart/pkg/auth"
	"github.com/gorilla/mux"
)

// Handlers handles users routes
func Handlers(r *mux.Router, service Service, userService users.Service) {
	r.Use(middleware.Authorize)
	r.Path("/items").Methods(http.MethodGet).HandlerFunc(RetrieveUserCart(service, userService))
	r.Path("/items").Methods(http.MethodPut).HandlerFunc(AddCartItem(service, userService))
}

// RetrieveUserCart retrieve User cart from DB
func RetrieveUserCart(service Service, userService users.Service) func(w http.ResponseWriter, r *http.Request) {
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

		cart, err := service.GetUserCart(ctx, user.ID)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusInternalServerError, "internal_error", err.Error())
		}
		httpresponse.RespondJSON(w, http.StatusOK, cart, nil)
	}
}

// addCartItemRequest htpp request for adding item in user cart
type addCartItemRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

func (req addCartItemRequest) Validate() error {
	var errorArray []string
	validator.CheckRule(&errorArray, req.ProductID > 0, "product_id is mandatory")
	validator.CheckRule(&errorArray, req.Quantity > 0, "quantity should be greater than 0")
	if len(errorArray) > 0 {
		return errors.New(strings.Join(errorArray, "; "))
	}
	return nil
}

// AddCartItem add item in user cart
func AddCartItem(service Service, userService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := loglib.GetLogger(ctx)
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

		logger.Infof("user is %v", user.Username)
		// unmarshal request
		req := addCartItemRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); (err != nil || req == addCartItemRequest{}) {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		// validate request
		if err := req.Validate(); err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		cart, err := service.AddItemCart(ctx, user.ID, req.ProductID, req.Quantity)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusInternalServerError, "internal_error", err.Error())
			return
		}

		httpresponse.RespondJSON(w, http.StatusOK, cart, nil)
	}
}
