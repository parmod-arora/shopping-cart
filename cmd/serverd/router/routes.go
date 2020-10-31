// Package router contains routing configuration for serverd
package router

import (
	"net/http"
	"os"

	"cinemo.com/shoping-cart/cmd/serverd/middleware"
	"cinemo.com/shoping-cart/internal/application"
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

	// // Register base middlewares
	// r.Use(middleware.Recover())
	r.Use(middleware.RequestLogger)
	r.StrictSlash(true)

	v1Router := r.PathPrefix("/api/v1").Subrouter()

	usersRouter := v1Router.PathPrefix("users").Subrouter()
	// // API routes
	// r.Group(api.Router(app))

	// // File server
	// r.Group(fs)

	return r
}
