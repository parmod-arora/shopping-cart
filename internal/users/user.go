package users

import "database/sql"

// Service is the interface to expose User functions
type Service interface {
	CreateUser(username string, password string, firstName *string, lastName *string) (*User, error)
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
func (s *userService) CreateUser(username string, password string, firstName *string, lastName *string) (*User, error) {
	return &User{
		Username: username,
	}, nil
}

func (s *userService) Validate(username string, password string) (*User, error) {
	return &User{
		Username: username,
	}, nil
}
