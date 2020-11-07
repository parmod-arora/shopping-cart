package discount

import (
	"context"
	"database/sql"
	"time"

	"cinemo.com/shoping-cart/internal/orm"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Service expose discount functions
type Service interface {
	RetrieveProductDiscounts(ctx context.Context, productID int64) ([]ProductDiscount, error)
}

type discountService struct {
	DB *sql.DB
}

// NewDiscountService create discount service
func NewDiscountService(db *sql.DB) Service {
	return &discountService{
		DB: db,
	}
}

// QuantityFunction expose eq|gte|lte functions
type QuantityFunction string

// EQ QuantityFunction
const EQ QuantityFunction = "EQ"

// GTE QuantityFunction
const GTE QuantityFunction = "GTE"

// LTE QuantityFunction
const LTE QuantityFunction = "LTE"

// Type discount
type Type string

// PERCENTAGE discount type
const PERCENTAGE Type = "PERCENTAGE"

// Rule discount rules
type Rule struct {
	ID                int64            `json:"id"`
	ProductDiscountID int64            `json:"product_discount_id"`
	ProductID         int64            `json:"product_id"`
	ProductQuantity   int64            `json:"product_quantity"`
	ProductQuantityFN QuantityFunction `json:"product_quantity_fn"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}

// ProductDiscount ProductDiscount
type ProductDiscount struct {
	ID           int64     `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Discount     int64     `json:"discount,omitempty"`
	DiscountType Type      `json:"discount_type,omitempty"`
	Rules        []Rule    `json:"discount_rules,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Quantity Quantity
type Quantity struct {
	Function QuantityFunction `json:"function,omitempty"`
	Value    int64            `json:"value,omitempty"`
}

// RetrieveProductDiscounts retrieve product discount
func (s *discountService) RetrieveProductDiscounts(ctx context.Context, productID int64) ([]ProductDiscount, error) {
	var discounts []ProductDiscount

	// find product offers for aproduct
	discountRules, err := orm.ProductDiscountRules(
		qm.Select("DISTINCT product_discount_rules.product_id, product_discount_rules.*"),
		qm.Where(orm.ProductDiscountRuleColumns.ProductID+"=?", productID),
	).All(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	for _, rule := range discountRules {
		ormProductDiscounts, err := orm.ProductDiscounts(
			qm.Load(qm.Rels(orm.ProductDiscountRels.ProductDiscountRules)),
			qm.Where(orm.ProductDiscountColumns.ID+"=?", rule.ProductDiscountID),
		).All(ctx, s.DB)
		if err != nil {
			return nil, err
		}

		for _, ormProductDiscount := range ormProductDiscounts {
			productDiscount := ProductDiscount{
				ID:           ormProductDiscount.ID,
				Name:         ormProductDiscount.Name,
				Discount:     ormProductDiscount.Discount,
				DiscountType: Type(ormProductDiscount.DiscountType),
				CreatedAt:    ormProductDiscount.CreatedAt,
				UpdatedAt:    ormProductDiscount.UpdatedAt,
			}

			for _, ormDiscountRule := range ormProductDiscount.R.ProductDiscountRules {
				rule := Rule{
					ID:                ormDiscountRule.ID,
					ProductDiscountID: ormDiscountRule.ProductDiscountID,
					ProductID:         ormDiscountRule.ProductID,
					ProductQuantity:   ormDiscountRule.ProductQuantity,
					ProductQuantityFN: QuantityFunction(ormDiscountRule.ProductQuantityFN),
					CreatedAt:         ormDiscountRule.CreatedAt,
					UpdatedAt:         ormDiscountRule.UpdatedAt,
				}
				productDiscount.Rules = append(productDiscount.Rules, rule)
			}
			discounts = append(discounts, productDiscount)
		}
	}
	return discounts, nil
}
