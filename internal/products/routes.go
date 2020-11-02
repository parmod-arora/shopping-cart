package products

import (
	"net/http"

	"cinemo.com/shoping-cart/framework/web/httpresponse"
	"cinemo.com/shoping-cart/framework/web/middleware"
	"github.com/gorilla/mux"
)

// Handlers handles users routes
func Handlers(r *mux.Router, service Service) {
	r.Use(middleware.Authorize)
	r.Path("/").Methods(http.MethodGet).HandlerFunc(ListProduct(service))
}

// ListProduct ListProduct
func ListProduct(service Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
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
