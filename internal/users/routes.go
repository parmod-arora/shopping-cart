package users

import (
	"encoding/json"
	"errors"
	"net/http"

	"cinemo.com/shoping-cart/framework/web/httpresponse"
	"cinemo.com/shoping-cart/internal/auth"
	"cinemo.com/shoping-cart/internal/errorcode"
	"github.com/gorilla/mux"
)

// Handlers handles users routes
func Handlers(r *mux.Router, service Service) {
	r.HandleFunc("/login", loginHandlers(service))
	r.HandleFunc("/signup", SignUpHandler(service))
}

func loginHandlers(service Service) func(http.ResponseWriter, *http.Request) {
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

// SignUpHandler SignUpHandler
func SignUpHandler(userService Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// unmarshal request
		req := userRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); (err != nil || req == userRequest{}) {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		// validate request
		if err := req.Validate(); err != nil {
			httpresponse.ErrorResponseJSON(ctx, w, http.StatusBadRequest, errorcode.ErrorsInRequestData, err.Error())
			return
		}

		// create user in database
		user, err := userService.CreateUser(ctx, req.Username, req.Password, req.FirstName, req.LastName)
		if err != nil {
			status, errCode := statusAndErrorCodeForServiceError(err)
			httpresponse.ErrorResponseJSON(ctx, w, status, errCode, err.Error())
			return
		}

		//prepare response
		response := &userResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		httpresponse.RespondJSON(w, http.StatusCreated, response, nil)
	}
}

func statusAndErrorCodeForServiceError(err error) (int, string) {
	if errors.As(err, &errorcode.ValidationError{}) {
		return http.StatusBadRequest, errorcode.ErrorsInRequestData
	} else if errors.As(err, &errorcode.DBError{}) {
		return http.StatusInternalServerError, errorcode.DatabaseProcessError
	}
	return http.StatusInternalServerError, errorcode.InternalError
}
