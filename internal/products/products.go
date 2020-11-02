package products

import "database/sql"

// Service is the interface to expose inventory functions
type Service interface {
	RetrieveProducts() ([]Product, error)
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

func (s *productService) RetrieveProducts() ([]Product, error) {
	return nil, nil
}
