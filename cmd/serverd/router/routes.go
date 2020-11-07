// Package router contains routing configuration for serverd
package router

import (
	"net/http"
	"os"

	"cinemo.com/shoping-cart/application"
	"cinemo.com/shoping-cart/framework/web/middleware"
	"cinemo.com/shoping-cart/internal/cart"
	"cinemo.com/shoping-cart/internal/products"
	"cinemo.com/shoping-cart/internal/users"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	wd string
)

func init() {
	var err error
	if wd, err = os.Getwd(); err != nil {
		panic(err.Error())
	}
}

// Handler returns the http handler that handles all requests
func Handler(app *application.App) http.Handler {
	r := mux.NewRouter()
	r.Use(handlers.RecoveryHandler())
	r.Use(middleware.RequestLogger)
	r.StrictSlash(true)
	v1 := r.PathPrefix("/api/v1").Subrouter()
	users.Handlers(v1.PathPrefix("/users").Subrouter(), app.UserService)
	products.Handlers(v1.PathPrefix("/products").Subrouter(), app.ProductService, app.UserService)
	cart.Handlers(v1.PathPrefix("/carts").Subrouter(), app.CartService, app.UserService)
	return r
}
