package products

import (
	"net/http"

	"cinemo.com/shoping-cart/framework/web/middleware"
	"github.com/gorilla/mux"
)

// Handlers handles users routes
func Handlers(r *mux.Router, service Service) {
	r.Use(middleware.Authorize)
	r.Path("/").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the book
		// navigate to the page
		w.Write([]byte("OKAY"))
		w.WriteHeader(200)
	})

}
