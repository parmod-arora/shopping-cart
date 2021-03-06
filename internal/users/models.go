package users

import (
	"errors"
	"strings"
	"time"

	"cinemo.com/shoping-cart/framework/web/validator"
)

// User structure in app
type User struct {
	ID        int64
	FirstName *string
	LastName  *string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserRequest is what we require from clients when adding a User.
type userRequest struct {
	Username  string  `json:"email" validate:"required"`
	Password  string  `json:"password" validate:"required"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type userResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"email"`
	FirstName *string   `json:"firstName,omitempty"`
	LastName  *string   `json:"lastName,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//Validate provides the validation rule for UserRequest
func (req *userRequest) Validate() error {
	var errorArray []string
	validator.CheckRule(&errorArray, len(req.Username) > 0, "username is mandatory")
	validator.CheckRule(&errorArray, len(req.Password) > 0, "password is mandatory")
	if len(errorArray) > 0 {
		return errors.New(strings.Join(errorArray, "; "))
	}
	return nil
}

type loginRequest struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (req *loginRequest) Validate() error {
	var errorArray []string
	validator.CheckRule(&errorArray, len(req.Username) > 0, "username is mandatory")
	validator.CheckRule(&errorArray, len(req.Password) > 0, "password is mandatory")
	if len(errorArray) > 0 {
		return errors.New(strings.Join(errorArray, "; "))
	}
	return nil
}
