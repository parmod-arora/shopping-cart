package coupons_test

import (
	"context"
	"database/sql"
	"reflect"
	"strconv"
	"testing"
	"time"

	"cinemo.com/shoping-cart/internal/coupons"
	"cinemo.com/shoping-cart/internal/discounts"
	"cinemo.com/shoping-cart/internal/products"
	"cinemo.com/shoping-cart/pkg/testutil"
	"cinemo.com/shoping-cart/pkg/trace"
)

func Test_couponService_CreateCoupon(t *testing.T) {
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
		db              *sql.DB
		productService  products.Service
		discountService discounts.Service
	}
	type args struct {
		ctx context.Context
	}
	now := time.Now().UTC()
	tests := []struct {
		name    string
		fixture string
		fields  fields
		args    args
		want    *coupons.Coupon
		wantErr bool
	}{
		{
			name: "add coupon in db",
			args: args{
				ctx: context.Background(),
			},
			fields: fields{
				db:              dbConnPool,
				discountService: discounts.NewDiscountService(dbConnPool),
				productService:  products.NewProductService(dbConnPool),
			},
			fixture: "testdata/add_coupon.sql",
			want: &coupons.Coupon{
				ID:         1,
				Name:       "ORANGE_" + strconv.FormatInt(now.UnixNano(), 10),
				ExpireAt:   now.Add(time.Hour),
				DiscountID: 3,
				ProductID:  4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.LoadFixture(dbConnPool, tt.fixture)
			s := coupons.NewCouponService(tt.fields.db, tt.fields.productService, tt.fields.discountService)

			got, err := s.CreateCoupon(tt.args.ctx, now)
			if (err != nil) != tt.wantErr {
				t.Errorf("couponService.CreateCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("couponService.CreateCoupon() = %v, want %v", got, tt.want)
			}
		})
	}
}
