package web

import (
	"net/http"

	"github.com/pkg/errors"
)

// Error represents a handler error. It contains web-related information such as
// HTTP status code, error code, description, and embeds the built-in error interface.
type Error struct {
	Status int    `json:"-"`
	Code   string `json:"error"`
	Desc   string `json:"error_description"`
	Err    error  `json:"-"`
}

func (e Error) Error() string {
	return e.Desc
}

// TypecastError performs a type assertion on the provide `error` and returns the object if concrete type is `Error`
func TypecastError(err error) *Error {
	if err == nil {
		return nil
	}

	if werr, ok := err.(*Error); ok {
		return werr
	}
	return nil
}

// WithStack adds stack trace into the Error object
func WithStack(err error) *Error {
	werr := TypecastError(err)
	if werr == nil {
		werr = &Error{Status: http.StatusInternalServerError, Code: "internal_error", Desc: err.Error()}
	}
	if werr.Err == nil {
		werr.Err = errors.WithStack(errors.New(werr.Error()))
	}
	return werr
}

// NewError returns a new Error object based on the provided err and message
func NewError(err error, message string) *Error {
	var result *Error
	werr := TypecastError(err)
	if werr != nil {
		result = &Error{
			Status: werr.Status,
			Code:   werr.Code,
			Desc:   message,
		}
	} else {
		result = &Error{Status: http.StatusInternalServerError, Code: "internal_error", Desc: message}
	}
	result.Err = errors.WithStack(errors.New(message))
	return result
}
