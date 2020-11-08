package products

import (
	"context"
	"database/sql"

	"cinemo.com/shoping-cart/internal/errorcode"
	"cinemo.com/shoping-cart/internal/orm"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (s *productService) RetrieveProductIDByName(ctx context.Context, name string) (int64, error) {
	product, err := orm.Products(qm.Where(orm.ProductColumns.Name+"=?", name)).One(ctx, s.DB)
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (s *productService) RetrieveProducts(ctx context.Context) ([]Product, error) {
	products, err := orm.Products().All(ctx, s.DB)
	if err != nil {
		return nil, errorcode.DBError{Err: err}
	}
	var models []Product
	for _, product := range products {
		models = append(models, *transformOrmToModelProduct(product))
	}
	return models, nil
}

func transformOrmToModelProduct(product *orm.Product) *Product {
	return &Product{
		ID:        product.ID,
		Amount:    product.Amount,
		Details:   product.Details,
		Name:      product.Name,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		Image:     product.Image,
	}
}
