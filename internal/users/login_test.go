package users_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"cinemo.com/shoping-cart/internal/users"
	mocks "cinemo.com/shoping-cart/mocks/users"
	"cinemo.com/shoping-cart/pkg/projectpath"
	"cinemo.com/shoping-cart/pkg/yaml"
	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

func TestLoginHandlers(t *testing.T) {

	// load jwt cert config
	vars, err := yaml.FetchEnvVarsFromYaml(projectpath.Root + "/jwt-cert.yml")
	if err != nil {
		logrus.Fatalf("Error while loading jwt cert %v", err.Error())
	}
	yaml.SetEnvVars(vars)

	tests := []struct {
		name                         string
		givenJSONReqFilepath         string
		givenUserServiceArgs         []interface{}
		givenUserServiceReturnValues []interface{}
		expectedJSONRespFilepath     string
		expectedStatusCode           int
	}{
		{
			name:                     "ideal success case",
			givenJSONReqFilepath:     "testdata/login/valid_request.json",
			expectedJSONRespFilepath: "testdata/login/success_response.json",
			expectedStatusCode:       http.StatusOK,
			givenUserServiceArgs:     []interface{}{mock.Anything, "email@email.com", "password"},
			givenUserServiceReturnValues: []interface{}{&users.User{
				ID:        1,
				Username:  "email@emalil.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil},
		},
		{
			name:                     "invalid request case",
			givenJSONReqFilepath:     "testdata/login/invalid_request.json",
			expectedJSONRespFilepath: "testdata/login/invalid_username_response.json",
			expectedStatusCode:       http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userService := new(mocks.Service)
			userService.On("Validate", tt.givenUserServiceArgs...).Return(tt.givenUserServiceReturnValues...)

			// Load test fixtures
			input, err := ioutil.ReadFile(tt.givenJSONReqFilepath)
			if err != nil {
				t.Fatalf("Cannot read %v", tt.givenJSONReqFilepath)
			}

			expected, err := ioutil.ReadFile(tt.expectedJSONRespFilepath)
			if err != nil {
				t.Fatalf("Cannot read %v", tt.expectedJSONRespFilepath)
			}

			r := httptest.NewRequest(http.MethodPost, "/v1/api/users/login", bytes.NewBuffer(input))
			w := httptest.NewRecorder()

			// when
			users.LoginHandlers((userService))(w, r)

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

			if tt.expectedStatusCode != 200 {
				if !cmp.Equal(expectedResult, result) {
					t.Errorf("expected response  diff: %v %v", result, cmp.Diff(expectedResult, result))
				}
			}
		})
	}
}
