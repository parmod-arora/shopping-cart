package orders_test

import (
	"context"
	"database/sql"
	"testing"

	"cinemo.com/shoping-cart/internal/discount"
	"cinemo.com/shoping-cart/internal/orders"
	"cinemo.com/shoping-cart/internal/products"
	"cinemo.com/shoping-cart/pkg/testutil"
	"cinemo.com/shoping-cart/pkg/trace"
	"github.com/google/go-cmp/cmp"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Test_orderService_GetUserCart(t *testing.T) {
	boil.DebugMode = true
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

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		fixture string
		fields  fields
		args    args
		wantErr bool
		want    *orders.UserCart
	}{
		{
			name:    "test sql query",
			fixture: "testdata/user_cart.sql",
			args: args{
				ctx:    context.Background(),
				userID: 2,
			},
			fields: fields{
				DB: dbConnPool,
			},
			want: &orders.UserCart{
				LineItems: []orders.LineItem{
					{
						CartItem: orders.CartItem{
							ID:       1,
							Product:  products.Product{ID: 1, Name: "Apples", Details: "Apples Details", Amount: 1000},
							Quantity: 10,
						},
						Discount: &discount.ProductDiscount{
							Name:         "10% Discount on 7+ Apples",
							Discount:     10,
							DiscountType: discount.PERCENTAGE,
							Quantity:     discount.Quantity{Function: discount.GTE, Value: 7},
						},
						Amount:           10000,
						DiscountedAmount: 900,
						SubTotal:         9000,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.LoadFixture(tt.fields.DB, tt.fixture)
			s := orders.NewOrderService(tt.fields.DB, discount.NewDiscountService(tt.fields.DB))
			got, err := s.GetUserCart(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderService.GetUserCart() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(got.LineItems, tt.want.LineItems) {
				t.Errorf("line items diff %v", cmp.Diff(got.LineItems, tt.want.LineItems))
			}
		})
	}
}
