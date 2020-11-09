package users

import (
	"errors"
	"net/http"

	"cinemo.com/shoping-cart/internal/errorcode"
	"github.com/gorilla/mux"
)

// Handlers handles users routes
func Handlers(r *mux.Router, service Service) {
	r.Path("/login").Methods(http.MethodPost).HandlerFunc(LoginHandlers(service))
	r.Path("/signup").Methods(http.MethodPost).HandlerFunc(SignUpHandler(service))
}

func statusAndErrorCodeForServiceError(err error) (int, string) {
	if errors.As(err, &errorcode.ValidationError{}) {
		return http.StatusBadRequest, errorcode.ErrorsInRequestData
	} else if errors.As(err, &errorcode.DBError{}) {
		return http.StatusInternalServerError, errorcode.DatabaseProcessError
	}
	return http.StatusInternalServerError, errorcode.InternalError
}
