package users_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "cinemo.com/shoping-cart/internal/users"
	mocks "cinemo.com/shoping-cart/mocks/users"
	"cinemo.com/shoping-cart/pkg/pointer"
	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func Test_signUpHandler(t *testing.T) {
	type args struct {
		userService Service
	}
	tests := []struct {
		name                         string
		givenJSONReqFilepath         string
		givenUserServiceArgs         []interface{}
		givenUserServiceReturnValues []interface{}
		expectedJSONRespFilepath     string
		expectedStatusCode           int
	}{
		{
			name:                     "ideal case success handler",
			givenJSONReqFilepath:     "testdata/signup/success/request.json",
			givenUserServiceArgs:     []interface{}{mock.Anything, "username", "password", pointer.String("firstname"), pointer.String("lastname")},
			expectedJSONRespFilepath: "testdata/signup/success/response.json",
			givenUserServiceReturnValues: []interface{}{&User{
				ID:        int64(1),
				FirstName: pointer.String("firstname"),
				LastName:  pointer.String("lastname"),
				Password:  "zxzx",
				Username:  "username",
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			}, nil},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:                     "ideal case success handler without first/lastname",
			givenJSONReqFilepath:     "testdata/signup/success/request_without_first_lastname.json",
			givenUserServiceArgs:     []interface{}{mock.Anything, "username", "password", mock.Anything, mock.Anything},
			expectedJSONRespFilepath: "testdata/signup/success/response_without_first_lastname.json",
			givenUserServiceReturnValues: []interface{}{&User{
				ID:        int64(1),
				Password:  "zxzx",
				Username:  "username",
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			}, nil},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:                         "invalid request",
			givenJSONReqFilepath:         "testdata/signup/failure/invalid_request",
			givenUserServiceArgs:         []interface{}{},
			expectedJSONRespFilepath:     "testdata/signup/failure/invalid_request_response.json",
			givenUserServiceReturnValues: []interface{}{},
			expectedStatusCode:           http.StatusBadRequest,
		},
		{
			name:                         "invalid request",
			givenJSONReqFilepath:         "testdata/signup/failure/request_without_username.json",
			givenUserServiceArgs:         []interface{}{},
			expectedJSONRespFilepath:     "testdata/signup/failure/response_of_request_without_username.json",
			givenUserServiceReturnValues: []interface{}{},
			expectedStatusCode:           http.StatusBadRequest,
		},
		{
			name:                         "invalid request",
			givenJSONReqFilepath:         "testdata/signup/failure/request_without_password.json",
			givenUserServiceArgs:         []interface{}{},
			expectedJSONRespFilepath:     "testdata/signup/failure/response_of_request_without_password.json",
			givenUserServiceReturnValues: []interface{}{},
			expectedStatusCode:           http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			userService := new(mocks.Service)
			userService.On("CreateUser", tt.givenUserServiceArgs...).Return(tt.givenUserServiceReturnValues...)

			// Load test fixtures
			input, err := ioutil.ReadFile(tt.givenJSONReqFilepath)
			if err != nil {
				t.Fatalf("Cannot read %v", tt.givenJSONReqFilepath)
			}

			expected, err := ioutil.ReadFile(tt.expectedJSONRespFilepath)
			if err != nil {
				t.Fatalf("Cannot read %v", tt.expectedJSONRespFilepath)
			}

			r := httptest.NewRequest(http.MethodPost, "/v1/api/users/singup", bytes.NewBuffer(input))
			w := httptest.NewRecorder()

			// when
			SignUpHandler(userService)(w, r)

			// then
			// prettify so we can match with expected file
			resp := w.Result()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("cannot read response: +%v", resp.Body)
			}

			var result interface{}
			err = json.Unmarshal(body, &result)
			if err != nil {
				t.Fatalf("cannot unmarshal response: %v", string(body))
			}

			var expectedResult interface{}
			err = json.Unmarshal(expected, &expectedResult)
			if err != nil {
				t.Fatalf("cannot unmarshal response: %v", string(expected))
			}

			// Assert:
			if !cmp.Equal(tt.expectedStatusCode, resp.StatusCode) {
				t.Errorf("status code mismatch diff: %v", cmp.Diff(tt.expectedStatusCode, resp.StatusCode))
			}

			if !cmp.Equal(expectedResult, result) {
				t.Errorf("expected response  diff: %v %v", result, cmp.Diff(expectedResult, result))
			}
		})
	}
}

func TestHandlers(t *testing.T) {
	type args struct {
		r       *mux.Router
		service Service
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "should run succesfully",
			args: args{
				r:       mux.NewRouter(),
				service: new(mocks.Service),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Handlers(tt.args.r, tt.args.service)
		})
	}
}
