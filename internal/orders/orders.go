package orders

import (
	"context"
	"database/sql"

	"cinemo.com/shoping-cart/internal/orm"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Service is the interface to expose order functions
type Service interface {
}

type orderService struct {
	DB *sql.DB
}

// NewOrderService Create New OrderService
func NewOrderService(db *sql.DB) Service {
	return orderService{
		DB: db,
	}
}

// AddItemCart add product in cart by user
func (s orderService) AddItemCart(ctx context.Context, userID int64, productID int64, quantity int64) error {
	orm.Orders(qm.Where()).All(ctx, s.DB)
	return nil
}
