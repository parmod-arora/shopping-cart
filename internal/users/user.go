package users

import (
	"context"
	"database/sql"
	"strings"

	"cinemo.com/shoping-cart/framework/loglib"
	"cinemo.com/shoping-cart/internal/errorcode"
	"cinemo.com/shoping-cart/internal/orm"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

// Service is the interface to expose User functions
type Service interface {
	CreateUser(ctx context.Context, username string, password string, firstName *string, lastName *string) (*User, error)
	Validate(username string, password string) (*User, error)
}

type userService struct {
	DB *sql.DB
}

// NewUserService method create a instance of userService
func NewUserService(db *sql.DB) Service {
	return &userService{
		DB: db,
	}
}

// CreateUser method creates a user in DB
func (s *userService) CreateUser(ctx context.Context, username string, password string, firstName *string, lastName *string) (*User, error) {
	logger := loglib.GetLogger(ctx)

	// Convert all username to lower before storing
	username = strings.ToLower(username)
	passwordHash, err := EncryptPassword(password)
	if err != nil {
		logger.Errorf("error: EncryptPassword %s", err.Error())
		return nil, err
	}

	existingUser, err := retrieveUserByUsername(ctx, s.DB, username)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("error: retrieveUserByUsername %s", err.Error())
			return nil, err
		}
	}

	if existingUser != nil {
		logger.Infof("User already exists with username %s", username)
		return nil, errorcode.ValidationError{Err: errors.Errorf("user already exists with %s", username)}
	}

	return saveUser(ctx, s.DB, &User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Password:  passwordHash,
	})
}

func retrieveUserByUsername(ctx context.Context, db *sql.DB, username string) (*User, error) {
	ormUser, err := orm.Users(
		qm.Where(orm.UserColumns.Username+"=?", username),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}
	return transformOrmToModelUser(ormUser), nil
}

func saveUser(ctx context.Context, db *sql.DB, user *User) (*User, error) {
	logger := loglib.GetLogger(ctx)
	ormUser := transformModelUserToOrm(user)
	if err := ormUser.Insert(ctx, db, boil.Infer()); err != nil {
		logger.Errorf("error: saveUser %s", err.Error())
		return nil, err
	}
	return transformOrmToModelUser(ormUser), nil
}

func transformModelUserToOrm(user *User) *orm.User {
	return &orm.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		Firstname: null.StringFromPtr(user.FirstName),
		Lastname:  null.StringFromPtr(user.LastName),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func transformOrmToModelUser(user *orm.User) *User {
	return &User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		FirstName: user.Firstname.Ptr(),
		LastName:  user.Lastname.Ptr(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (s *userService) Validate(username string, password string) (*User, error) {
	return &User{
		Username: username,
	}, nil
}

// EncryptPassword generates hashed password from string
func EncryptPassword(clearPassword string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(clearPassword), 12)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}
