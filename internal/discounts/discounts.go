package discounts

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
	FetchDiscountIDByName(ctx context.Context, name string) (int64, error)
	FetchDiscountByID(ctx context.Context, id int64) (ProductDiscount, error)
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

func (s *discountService) FetchDiscountByID(ctx context.Context, ID int64) (ProductDiscount, error) {
	var productDiscount ProductDiscount
	ormProductDiscount, err := orm.Discounts(
		qm.Load(qm.Rels(orm.DiscountRels.DiscountRules)),
		qm.Where(orm.DiscountColumns.ID+"=?", ID),
	).One(ctx, s.DB)
	if err != nil {
		return productDiscount, err
	}
	return transformOrmToModel(ormProductDiscount), nil
}

func transformOrmToModel(ormProductDiscount *orm.Discount) ProductDiscount {
	var productDiscount ProductDiscount
	productDiscount = ProductDiscount{
		ID:           ormProductDiscount.ID,
		Name:         ormProductDiscount.Name,
		Discount:     ormProductDiscount.Discount,
		DiscountType: Type(ormProductDiscount.DiscountType),
		CreatedAt:    ormProductDiscount.CreatedAt,
		UpdatedAt:    ormProductDiscount.UpdatedAt,
	}

	for _, ormDiscountRule := range ormProductDiscount.R.DiscountRules {
		rule := Rule{
			ID:                ormDiscountRule.ID,
			ProductDiscountID: ormDiscountRule.DiscountID,
			ProductID:         ormDiscountRule.ProductID,
			ProductQuantity:   ormDiscountRule.ProductQuantity,
			ProductQuantityFN: QuantityFunction(ormDiscountRule.ProductQuantityFN),
			CreatedAt:         ormDiscountRule.CreatedAt,
			UpdatedAt:         ormDiscountRule.UpdatedAt,
		}
		productDiscount.Rules = append(productDiscount.Rules, rule)
	}
	return productDiscount
}

// RetrieveProductDiscounts retrieve product discount
func (s *discountService) RetrieveProductDiscounts(ctx context.Context, productID int64) ([]ProductDiscount, error) {
	var discounts []ProductDiscount
	// find product offers for aproduct
	discountRules, err := orm.DiscountRules(
		qm.Select("DISTINCT discount_rules.product_id, discount_rules.*"),
		qm.Where(orm.DiscountRuleColumns.ProductID+"=?", productID),
	).All(ctx, s.DB)
	if err != nil {
		return nil, err
	}
	var discountIds []interface{}
	for _, rule := range discountRules {
		discountIds = append(discountIds, rule.DiscountID)
	}

	ormDiscounts, err := orm.Discounts(
		qm.Load(qm.Rels(orm.DiscountRels.DiscountRules)),
		qm.LeftOuterJoin(orm.TableNames.Coupons+" on "+orm.CouponColumns.DiscountID+"="+orm.TableNames.Discounts+"."+orm.DiscountColumns.ID),
		qm.WhereIn(orm.TableNames.Discounts+"."+orm.DiscountColumns.ID+" in ?", discountIds...),
		qm.And(orm.TableNames.Coupons+"."+orm.CouponColumns.ID+" is null"),
	).All(ctx, s.DB)
	if err != nil {
		return discounts, err
	}

	for _, ormDiscount := range ormDiscounts {
		discounts = append(discounts, transformOrmToModel(ormDiscount))
	}
	return discounts, nil
}

func (s *discountService) FetchDiscountIDByName(ctx context.Context, name string) (int64, error) {
	discount, err := orm.Discounts(qm.Where(orm.DiscountColumns.Name+"=?", name)).One(ctx, s.DB)
	if err != nil {
		return 0, err
	}
	return discount.ID, nil
}
