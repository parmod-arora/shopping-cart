package cart_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"cinemo.com/shoping-cart/internal/cart"
	"cinemo.com/shoping-cart/internal/coupons"
	"cinemo.com/shoping-cart/internal/discounts"
	"cinemo.com/shoping-cart/internal/products"
	"cinemo.com/shoping-cart/pkg/testutil"
	"cinemo.com/shoping-cart/pkg/trace"
	"github.com/google/go-cmp/cmp"
)

func Test_orderService_GetUserCart(t *testing.T) {
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
		want    *cart.UserCart
	}{
		{
			name:    "get user cart",
			fixture: "testdata/user_cart.sql",
			args: args{
				ctx:    context.Background(),
				userID: 2,
			},
			fields: fields{
				DB: dbConnPool,
			},
			want: &cart.UserCart{
				ID:                2,
				UserID:            2,
				SubTotalAmount:    11800,
				TotalSavingAmount: 1760,
				TotalAmount:       10040,
				CartItems: []cart.CartItem{
					{
						ID:     2,
						CartID: 2,
						Product: products.Product{
							ID:        1,
							Name:      "Apples",
							Details:   "Apples Details",
							Amount:    1000,
							Image:     "apple.jpeg",
							CreatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
							UpdatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
						},
						Quantity: 8,
						SubTotal: 8000,
					},
					{
						ID:     3,
						CartID: 2,
						Product: products.Product{
							ID:        2,
							Name:      "Bananas",
							Details:   "Bananas Details",
							Amount:    200,
							Image:     "banana.jpg",
							CreatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
							UpdatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
						},
						Quantity: 5,
						SubTotal: 1000,
					},
					{
						ID:     4,
						CartID: 2,
						Product: products.Product{
							ID:        3,
							Name:      "Pears",
							Details:   "Pears Details",
							Amount:    300,
							Image:     "pears.jpg",
							CreatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
							UpdatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
						},
						Quantity: 9,
						SubTotal: 2700,
					},
					{
						ID:     5,
						CartID: 2,
						Product: products.Product{
							ID:        4,
							Name:      "Oranges",
							Details:   "Oranges Details",
							Amount:    100,
							Image:     "orange.jpeg",
							CreatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
							UpdatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
						},
						Quantity: 1,
						SubTotal: 100,
					},
				},
				Coupon: &coupons.Coupon{
					ID:         1,
					Name:       "COUPON_30",
					ExpireAt:   time.Date(2020, 11, 8, 7, 22, 54, 0, time.UTC),
					RedeemedAt: nil,
					IsExpired:  true,
					ProductID:  4,
					DiscountID: 3,
				},
				LineItems: []cart.LineItem{
					{
						ProductDiscount: &discounts.ProductDiscount{
							ID:           1,
							Name:         "Apple 10 discount on 7 or more Apples",
							Discount:     10,
							DiscountType: "PERCENTAGE",
							Rules: []discounts.Rule{
								{
									ID:                1,
									ProductDiscountID: 1,
									ProductID:         1,
									ProductQuantity:   7,
									ProductQuantityFN: "GTE",
									CreatedAt:         time.Date(2020, 11, 7, 7, 16, 18, 0, time.UTC),
									UpdatedAt:         time.Date(2020, 11, 7, 7, 16, 18, 0, time.UTC),
								},
							},
							CreatedAt: time.Date(2020, 11, 7, 7, 5, 50, 0, time.UTC),
							UpdatedAt: time.Date(2020, 11, 7, 7, 5, 50, 0, time.UTC),
						},
						CartItems: []cart.CartItem{
							{
								ID:     2,
								CartID: 2,
								Product: products.Product{
									ID:        1,
									Name:      "Apples",
									Details:   "Apples Details",
									Amount:    1000,
									Image:     "apple.jpeg",
									CreatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
									UpdatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
								},
								Quantity: 8,
								SubTotal: 8000,
							},
						},
						DiscountAmount: 800,
						Quantity:       1,
					},
					{
						ProductDiscount: &discounts.ProductDiscount{
							ID:           2,
							Name:         "Combo discount on 4Pears and 2 Banana",
							Discount:     30,
							DiscountType: "PERCENTAGE",
							CreatedAt:    time.Date(2020, 11, 7, 7, 13, 59, 0, time.UTC),
							UpdatedAt:    time.Date(2020, 11, 7, 7, 13, 59, 0, time.UTC),
							Rules: []discounts.Rule{
								{
									ID:                2,
									ProductDiscountID: 2,
									ProductID:         3,
									ProductQuantity:   4,
									ProductQuantityFN: "EQ",
									CreatedAt:         time.Date(2020, 11, 7, 7, 16, 40, 0, time.UTC),
									UpdatedAt:         time.Date(2020, 11, 7, 7, 16, 40, 0, time.UTC),
								},
								{
									ID:                3,
									ProductDiscountID: 2,
									ProductID:         2,
									ProductQuantity:   2,
									ProductQuantityFN: "EQ",
									CreatedAt:         time.Date(2020, 11, 7, 7, 17, 07, 0, time.UTC),
									UpdatedAt:         time.Date(2020, 11, 7, 7, 17, 07, 0, time.UTC),
								},
							},
						},
						CartItems: []cart.CartItem{
							{
								ID:     4,
								CartID: 2,
								Product: products.Product{
									ID:        3,
									Name:      "Pears",
									Details:   "Pears Details",
									Amount:    300,
									Image:     "pears.jpg",
									CreatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
									UpdatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
								},
								Quantity: 9,
								SubTotal: 2700,
							},
							{
								ID:     3,
								CartID: 2,
								Product: products.Product{
									ID:        2,
									Name:      "Bananas",
									Details:   "Bananas Details",
									Amount:    200,
									Image:     "banana.jpg",
									CreatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
									UpdatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
								},
								Quantity: 5,
								SubTotal: 1000,
							},
						},
						DiscountAmount: 960,
						Quantity:       2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.LoadFixture(tt.fields.DB, tt.fixture)
			if err != nil {
				t.Errorf("seed error %v", err.Error())
			}
			discountsService := discounts.NewDiscountService(tt.fields.DB)
			productService := products.NewProductService(tt.fields.DB)
			couponService := coupons.NewCouponService(tt.fields.DB, productService, discountsService)
			s := cart.NewCartService(tt.fields.DB, discountsService, couponService)
			got, err := s.GetUserCart(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderService.GetUserCart() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("line items diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func Test_cartService_AddItemCart(t *testing.T) {
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
		DB              *sql.DB
		DiscountService discounts.Service
	}
	type args struct {
		ctx       context.Context
		userID    int64
		productID int64
		quantity  int64
	}
	tests := []struct {
		name    string
		fixture string
		fields  fields
		args    args
		want    *cart.UserCart
		wantErr bool
	}{
		{
			name:    "should run succssfully",
			fixture: "testdata/add_cart.sql",
			fields: fields{
				DB:              dbConnPool,
				DiscountService: discounts.NewDiscountService(dbConnPool),
			},
			args: args{
				ctx:       context.Background(),
				productID: 1,
				quantity:  15,
				userID:    1,
			},
			want: &cart.UserCart{
				ID:                1,
				UserID:            1,
				SubTotalAmount:    15000,
				TotalSavingAmount: 0,
				TotalAmount:       15000,
				Coupon:            nil,
				CartItems: []cart.CartItem{
					{
						ID:     1,
						CartID: 1,
						Product: products.Product{
							ID:        1,
							Name:      "Apples",
							Details:   "Apples Details",
							Amount:    1000,
							Image:     "apple.jpeg",
							CreatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
							UpdatedAt: time.Date(2020, 11, 2, 7, 52, 39, 0, time.UTC),
						},
						Quantity: 15,
						SubTotal: 15000,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.LoadFixture(dbConnPool, tt.fixture)
			if err != nil {
				t.Errorf("error = %v", err.Error())
				return
			}
			discountService := discounts.NewDiscountService(dbConnPool)
			couponService := coupons.NewCouponService(dbConnPool, products.NewProductService(dbConnPool), discountService)
			s := cart.NewCartService(dbConnPool, discountService, couponService)
			got, err := s.AddItemCart(tt.args.ctx, tt.args.userID, tt.args.productID, tt.args.quantity)
			if (err != nil) != tt.wantErr {
				t.Errorf("cartService.AddItemCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("cartService.AddItemCart() = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}
