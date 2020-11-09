package products

import (
	"context"
	"database/sql"
)

// Service is the interface to expose inventory functions
type Service interface {
	RetrieveProducts(ctx context.Context) ([]Product, error)
	RetrieveProductIDByName(ctx context.Context, name string) (int64, error)
}

type productService struct {
	DB *sql.DB
}

// NewProductService create new nstance of Product Service
func NewProductService(db *sql.DB) Service {
	return &productService{
		DB: db,
	}
}
