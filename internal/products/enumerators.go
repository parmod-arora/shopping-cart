package products

// PriceType is enumerates price type category
type PriceType string

const (
	// Discount is discount on price percentage
	Discount PriceType = "DISCOUNT"
	// Quantity is discount on quantity
	Quantity PriceType = "Quantity"
)

// String is a toString method for enums
func (s PriceType) String() string {
	return string(s)
}
