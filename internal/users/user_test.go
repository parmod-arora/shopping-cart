package users

import (
	"context"
	"testing"

	"cinemo.com/shoping-cart/internal/testutil"
	"cinemo.com/shoping-cart/pkg/trace"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_userService_CreateUser(t *testing.T) {
	t.Parallel()
	// Start adding unique database schema, make the parallel testing possible
	// Create unique schema
	traceInfo := trace.Trace()
	dbConnPool, schema, err := testutil.PrepareDatabase(traceInfo)
	if err != nil {
		t.Fatalf("Not able to create unique schema dbConnPool: %v", err)
	}
	defer func() {
		dbConnPool.Exec("DROP SCHEMA  IF EXISTS " + schema + " CASCADE")
		dbConnPool.Close()
	}()

	type args struct {
		ctx       context.Context
		username  string
		password  string
		firstName *string
		lastName  *string
	}

	tests := []struct {
		name    string
		fixture string
		args    args
		want    *User
		wantErr bool
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				DB: dbConnPool,
			}
			got, err := s.CreateUser(tt.args.ctx, tt.args.username, tt.args.password, tt.args.firstName, tt.args.lastName)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(tt.want, got, cmpopts.IgnoreFields(User{}, "CreatedAt", "UpdatedAt")) {
				t.Errorf("userService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
