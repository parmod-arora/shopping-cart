package discounts_test

import (
	"context"
	"log"
	"testing"

	discount "cinemo.com/shoping-cart/internal/discounts"
	"cinemo.com/shoping-cart/pkg/testutil"
	"cinemo.com/shoping-cart/pkg/trace"
)

func Test_discountService_RetrieveProductDiscounts(t *testing.T) {
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
		fixture string
	}
	type args struct {
		ctx       context.Context
		productID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ideal case for product discount",
			fields: fields{
				fixture: "testdata/product_discount.sql",
			},
			args: args{
				ctx:       context.Background(),
				productID: 1,
			},
		},
		{
			name: "ideal case for product discount",
			args: args{
				ctx:       context.Background(),
				productID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testutil.LoadFixture(dbConnPool, tt.fields.fixture); err != nil {
				log.Fatalf("error %v", err.Error())
			}
			s := discount.NewDiscountService(dbConnPool)
			if _, err := s.RetrieveProductDiscounts(tt.args.ctx, tt.args.productID); (err != nil) != tt.wantErr {
				t.Errorf("discountService.RetrieveProductDiscounts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
