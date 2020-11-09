package users

import (
	"encoding/json"
	"net/http"

	"cinemo.com/shoping-cart/framework/web/httpresponse"
	"cinemo.com/shoping-cart/internal/errorcode"
	"cinemo.com/shoping-cart/pkg/auth"
)

// LoginHandlers handles login functionality
func LoginHandlers(service Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// unmarshal request
		req := loginRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); (err != nil || req == loginRequest{}) {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		// validate request
		if err := req.Validate(); err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		// validate User
		user, err := service.Validate(ctx, req.Username, req.Password)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusForbidden, errorcode.LoginFailed, err.Error())
			return
		}

		// create jwt token
		token, err := auth.CreateJWT(user.Username)
		if err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusInternalServerError, errorcode.CreateTokenFailed, err.Error())
			return
		}

		httpresponse.RespondJSON(w, http.StatusOK, loginResponse{
			Token: string(token),
		}, nil)
	}
}
