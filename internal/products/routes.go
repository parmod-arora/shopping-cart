package products

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handlers handles users routes
func Handlers(r *mux.Router, service Service) {
	r.Path("/").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the book
		// navigate to the page
		w.Write([]byte("OKAY"))
		w.WriteHeader(200)
	})

}
