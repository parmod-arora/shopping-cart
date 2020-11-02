package products_test

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	. "cinemo.com/shoping-cart/internal/products"
	"cinemo.com/shoping-cart/pkg/testutil"
	"cinemo.com/shoping-cart/pkg/trace"
	"github.com/google/go-cmp/cmp"
)

func Test_productService_RetrieveProducts(t *testing.T) {
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
		DB      *sql.DB
		fixture string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Product
		wantErr error
	}{
		{
			name: "ideal case success",
			args: args{
				ctx: context.Background(),
			},
			fields: fields{
				DB:      dbConnPool,
				fixture: "testdata/retrieve_products.sql",
			},
			want: []Product{
				{
					ID:        1,
					Name:      "Apples",
					Details:   "Apples Details",
					Amount:    1000,
					CreatedAt: time.Date(2020, 11, 02, 07, 52, 39, 0, time.UTC),
					UpdatedAt: time.Date(2020, 11, 02, 07, 52, 39, 0, time.UTC),
				},
				{
					ID:        2,
					Name:      "Bananas",
					Details:   "Bananas Details",
					Amount:    200,
					CreatedAt: time.Date(2020, 11, 02, 07, 52, 39, 0, time.UTC),
					UpdatedAt: time.Date(2020, 11, 02, 07, 52, 39, 0, time.UTC),
				},
				{
					ID:        3,
					Name:      "Pears",
					Details:   "Pears Details",
					Amount:    300,
					CreatedAt: time.Date(2020, 11, 02, 07, 52, 39, 0, time.UTC),
					UpdatedAt: time.Date(2020, 11, 02, 07, 52, 39, 0, time.UTC),
				},
				{
					ID:        4,
					Name:      "Oranges",
					Details:   "Oranges Details",
					Amount:    100,
					CreatedAt: time.Date(2020, 11, 02, 07, 52, 39, 0, time.UTC), // "2020-11-02 07:47:39.894527 +0000 UTC",
					UpdatedAt: time.Date(2020, 11, 02, 07, 52, 39, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testutil.LoadFixture(dbConnPool, tt.fields.fixture); err != nil {
				log.Fatalf("error %v", err.Error())
			}

			s := NewProductService(dbConnPool)
			got, err := s.RetrieveProducts(tt.args.ctx)
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("productService error is nil")
					return
				}
				if !cmp.Equal(err.Error(), tt.wantErr.Error()) {
					t.Errorf("productService want error = %v", cmp.Diff(got, tt.want))
				}
				return
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("productService.RetrieveProducts() = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}
