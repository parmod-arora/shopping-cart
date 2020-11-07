package cart_test

import (
	"context"
	"database/sql"
	"testing"

	"cinemo.com/shoping-cart/internal/cart"
	"cinemo.com/shoping-cart/internal/discount"
	"cinemo.com/shoping-cart/pkg/testutil"
	"cinemo.com/shoping-cart/pkg/trace"
	"github.com/google/go-cmp/cmp"
)

func Test_orderService_GetUserCart(t *testing.T) {
	t.Parallel()
	// Start adding unique database schema, make the parallel testing possible
	// Create unique schema
	traceInfo := trace.Trace()
	dbConnPool, _, err := testutil.PrepareDatabase(traceInfo)
	if err != nil {
		t.Fatalf("Not able to create unique schema dbConnPool: %v", err)
	}
	defer func() {
		// dbConnPool.Exec("DROP SCHEMA  IF EXISTS " + schema + " CASCADE")
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
		want    *cart.UserCart
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
			want: &cart.UserCart{
				LineItems: []cart.LineItem{
					{
						ProductDiscount: &discount.ProductDiscount{
							ID:           1,
							Name:         "Apple 10 % discount on 7 or more Apples",
							Discount:     10,
							DiscountType: "PERCENTAGE",
						},
						CartItems:      []cart.CartItem{{ID: 5, Quantity: 1}},
						DiscountAmount: 100,
						Quantity:       1,
					},
					{
						ProductDiscount: &discount.ProductDiscount{
							ID:           2,
							Name:         "Combo discount on 4Pears and 2 Banana",
							Discount:     30,
							DiscountType: "PERCENTAGE",
						},
						CartItems:      []cart.CartItem{},
						DiscountAmount: 150,
						Quantity:       2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.LoadFixture(tt.fields.DB, tt.fixture)
			s := cart.NewCartService(tt.fields.DB, discount.NewDiscountService(tt.fields.DB))
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

// func Test_cartService_AddItemCart(t *testing.T) {
// 	boil.DebugMode = true
// 	t.Parallel()
// 	// Start adding unique database schema, make the parallel testing possible
// 	// Create unique schema
// 	traceInfo := trace.Trace()
// 	dbConnPool, _, err := testutil.PrepareDatabase(traceInfo)
// 	if err != nil {
// 		t.Fatalf("Not able to create unique schema dbConnPool: %v", err)
// 	}
// 	defer func() {
// 		// dbConnPool.Exec("DROP SCHEMA  IF EXISTS " + schema + " CASCADE")
// 		dbConnPool.Close()
// 	}()

// 	type fields struct {
// 		DB              *sql.DB
// 		DiscountService discount.Service
// 	}
// 	type args struct {
// 		ctx       context.Context
// 		userID    int64
// 		productID int64
// 		quantity  int64
// 	}
// 	tests := []struct {
// 		name    string
// 		fixture string
// 		fields  fields
// 		args    args
// 		want    *cart.UserCart
// 		wantErr bool
// 	}{
// 		{
// 			name:    "should run succssfully",
// 			fixture: "testdata/add_cart.sql",
// 			fields: fields{
// 				DB:              dbConnPool,
// 				DiscountService: discount.NewDiscountService(dbConnPool),
// 			},
// 			args: args{
// 				ctx:       context.Background(),
// 				productID: 1,
// 				quantity:  15,
// 				userID:    1,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			testutil.LoadFixture(dbConnPool, tt.fixture)
// 			s := cart.NewCartService(dbConnPool, discount.NewDiscountService(dbConnPool))
// 			got, err := s.AddItemCart(tt.args.ctx, tt.args.userID, tt.args.productID, tt.args.quantity)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("cartService.AddItemCart() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("cartService.AddItemCart() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
