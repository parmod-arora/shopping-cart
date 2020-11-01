package httpresponse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type httpErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// RespondJSON writes JSON as http response
func RespondJSON(w http.ResponseWriter, httpStatusCode int, object interface{}, headers map[string]string) {
	bytes, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if headers != nil {
		for key, value := range headers {
			w.Header().Set(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatusCode)
	w.Write(bytes)
}

// RespondText writes Text as http response
func RespondText(w http.ResponseWriter, httpStatusCode int, text, fileName string, headers map[string]string) {
	if headers != nil {
		for key, value := range headers {
			w.Header().Set(key, value)
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	if fileName != "" {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%v.txt", fileName))
	}
	w.WriteHeader(httpStatusCode)
	w.Write([]byte(text))
}

// ErrorResponseJSON write ErrorResponse as http response
func ErrorResponseJSON(ctx context.Context, w http.ResponseWriter, httpStatusCode int, err string, description string) {

	// logger.Errorf("Error: %v, Response Code: %v, Description: %v", err, httpStatusCode, description)
	RespondJSON(w, httpStatusCode, httpErrorResponse{
		Error:            err,
		ErrorDescription: description,
	}, nil)
}
