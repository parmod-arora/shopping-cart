package discount

import (
	"context"
	"database/sql"

	"cinemo.com/shoping-cart/internal/orm"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Service expose discount functions
type Service interface {
	RetrieveProductDiscounts(ctx context.Context, productID int64) (map[int64]Rules, error)
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

// Rules discount rules
type Rules struct {
	ProductDiscounts []ProductDiscount              `json:"product_discounts,omitempty"`
	ComboDiscounts   map[int64]ComboDiscountProduct `json:"combo_discount,omitempty"`
}

// ComboDiscountProduct ComboDiscountProduct
type ComboDiscountProduct struct {
	Name                 string   `json:"name,omitempty"`
	PackagedWithQuantity Quantity `json:"packaged_with_quantity,omitempty"`
	Quantity             Quantity `json:"quantity,omitempty"`
	Discount             int64    `json:"disount,omitempty"`
	DiscountType         Type     `json:"disount_type,omitempty"`
}

// ProductDiscount ProductDiscount
type ProductDiscount struct {
	Name         string   `json:"name,omitempty"`
	Discount     int64    `json:"disount,omitempty"`
	DiscountType Type     `json:"disount_type,omitempty"`
	Quantity     Quantity `json:"quantity,omitempty"`
}

// Quantity Quantity
type Quantity struct {
	Function QuantityFunction `json:"function,omitempty"`
	Value    int64            `json:"value,omitempty"`
}

// RetrieveProductDiscounts retrieve product discount
func (s *discountService) RetrieveProductDiscounts(ctx context.Context, productID int64) (map[int64]Rules, error) {
	productRule := make(map[int64]Rules, 0)
	rules := Rules{}
	// retrieve product discount
	discounts, err := orm.ProductDiscounts(qm.Where(orm.ProductDiscountColumns.ProductID+"=?", productID)).All(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// prepare product discount
	for _, discount := range discounts {
		productDiscount := ProductDiscount{
			Name:         discount.Name,
			Discount:     discount.Discount,
			DiscountType: Type(discount.DiscountType),
			Quantity: Quantity{
				Function: QuantityFunction(discount.QuantityFN),
				Value:    discount.Quantity,
			},
		}
		rules.ProductDiscounts = append(rules.ProductDiscounts, productDiscount)
	}

	// retrieve combo product discount
	comboDiscounts, err := orm.ProductComboDiscounts(
		orm.ProductComboDiscountWhere.ProductID.EQ(productID),
		qm.Or(orm.ProductComboDiscountColumns.PackagedWithProductID+"=?", productID),
	).All(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// prepare combo discount
	comboDiscountMap := make(map[int64]ComboDiscountProduct)
	for _, discount := range comboDiscounts {
		// reverse entry
		packagedWithProductID := discount.ProductID
		packagedWithProductQuantity := discount.ProductQuantity
		productQuantity := discount.PackagedWithProductQuantity
		if packagedWithProductID == productID {
			packagedWithProductID = discount.PackagedWithProductID
			packagedWithProductQuantity = discount.PackagedWithProductQuantity
			productQuantity = discount.ProductQuantity
		}

		comboDiscount := ComboDiscountProduct{
			Name:         discount.Name,
			Discount:     discount.Discount,
			DiscountType: PERCENTAGE,
			PackagedWithQuantity: Quantity{
				Function: QuantityFunction(discount.PackagedWithProductQuantityFN),
				Value:    packagedWithProductQuantity,
			},
			Quantity: Quantity{
				Function: QuantityFunction(discount.ProductQuantityFN),
				Value:    productQuantity,
			},
		}
		comboDiscountMap[packagedWithProductID] = comboDiscount
	}
	rules.ComboDiscounts = comboDiscountMap
	productRule[productID] = rules
	return productRule, nil
}
