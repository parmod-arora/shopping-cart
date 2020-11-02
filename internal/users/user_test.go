package users_test

import (
	"context"
	"testing"

	"cinemo.com/shoping-cart/pkg/pointer"
	"cinemo.com/shoping-cart/pkg/testutil"
	"cinemo.com/shoping-cart/pkg/trace"
	"github.com/friendsofgo/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/crypto/bcrypt"

	. "cinemo.com/shoping-cart/internal/users"
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

	hashPassword, _ := EncryptPassword("password")
	tests := []struct {
		name    string
		fixture string
		args    args
		want    *User
		wantErr error
	}{
		{
			name: "Add new user to db",
			args: args{
				ctx:       context.Background(),
				firstName: pointer.String("firstname"),
				lastName:  pointer.String("lastname"),
				password:  "password",
				username:  "username",
			},
			want: &User{
				ID:        int64(1),
				FirstName: pointer.String("firstname"),
				LastName:  pointer.String("lastname"),
				Password:  hashPassword,
				Username:  "username",
			},
		},
		{
			name:    "Add duplicate user to db",
			fixture: "testdata/duplicate_user.sql",
			args: args{
				ctx:      context.Background(),
				username: "username",
				password: "password",
			},
			wantErr: errors.Errorf("user already exists with username"),
		},
		{
			name: "Add user without firstname",
			args: args{
				ctx:      context.Background(),
				lastName: pointer.String("lastname"),
				password: "password",
				username: "user_without_first_name",
			},
			want: &User{
				ID:       int64(2),
				LastName: pointer.String("lastname"),
				Password: hashPassword,
				Username: "user_without_first_name",
			},
		},
		{
			name: "Add user without lastname",
			args: args{
				ctx:       context.Background(),
				firstName: pointer.String("firstname"),
				password:  "password",
				username:  "user_without_lastname",
			},
			want: &User{
				ID:        int64(3),
				FirstName: pointer.String("firstname"),
				Password:  hashPassword,
				Username:  "user_without_lastname",
			},
		},
		{
			name: "Add user without firstname/lastname",
			args: args{
				ctx:      context.Background(),
				password: "password",
				username: "user_without_firstname_lastname",
			},
			want: &User{
				ID:       int64(4),
				Password: hashPassword,
				Username: "user_without_firstname_lastname",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.LoadFixture(dbConnPool, tt.fixture)
			s := NewUserService(dbConnPool)
			got, err := s.CreateUser(tt.args.ctx, tt.args.username, tt.args.password, tt.args.firstName, tt.args.lastName)
			if tt.wantErr != nil {
				if !cmp.Equal(tt.wantErr.Error(), err.Error()) {
					t.Errorf("userService.CreateUser() error diff = %v", cmp.Diff(tt.wantErr.Error(), err.Error()))
				}
				return
			}

			if !cmp.Equal(tt.want, got, cmpopts.IgnoreFields(User{}, "CreatedAt", "UpdatedAt", "Password")) {
				t.Errorf("userService.CreateUser() diff = %v", cmp.Diff(tt.want, got, cmpopts.IgnoreFields(User{}, "CreatedAt", "UpdatedAt", "Password")))
			}
			// password should match
			if err = bcrypt.CompareHashAndPassword([]byte(got.Password), []byte(tt.args.password)); err != nil {
				t.Errorf("userService.CreateUser() password mismatch  = %v ", err.Error())
			}
			// createdAt and updatedAT should not be nil
			if got.CreatedAt.IsZero() || got.UpdatedAt.IsZero() {
				t.Errorf("userService.CreateUser() created or updated is zero ")
			}
		})
	}
}
