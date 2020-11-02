package middleware

import (
	"net/http"

	"cinemo.com/shoping-cart/framework/web/httpresponse"
	"cinemo.com/shoping-cart/internal/errorcode"
	"cinemo.com/shoping-cart/pkg/auth"
	"github.com/SermoDigital/jose/jws"
)

// Authorize authorize user
func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		j, err := jws.ParseJWTFromRequest(r)
		if err != nil {
			httpresponse.ErrorResponseJSON(r.Context(), w, http.StatusUnauthorized, errorcode.AuthHeaderReadFailed, err.Error())
			return
		}

		err = auth.ValidateJWT(j)
		if err != nil {
			httpresponse.ErrorResponseJSON(r.Context(), w, http.StatusUnauthorized, errorcode.AuthValidateFailed, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
