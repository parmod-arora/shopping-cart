package users_test

import (
	"testing"

	. "cinemo.com/shoping-cart/internal/users"
	mocks "cinemo.com/shoping-cart/mocks/users"
	"github.com/gorilla/mux"
)

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
