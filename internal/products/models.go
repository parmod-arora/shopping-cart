package products

import "time"

// Product representation in app
type Product struct {
	ID        int64
	Name      string
	Details   string
	Amount    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductDiscount struct {
	ID             int64
	ProductID      int64
	MinQuantity    int
	MaxQuantity    int
	PriceType      PriceType
	ComboPackageID int64
	Discount       int
	Active         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Stock representation in app
type Stock struct {
	ID        int64
	ProductID int64
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Order representation in app
type Order struct {
	ID        int64
	Name      string
	Details   string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ComboPackage struct {
	ID           int64
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ComboPackagedWith struct {
	ID                    int64
	ProductID             int64
	PackagedWithProductID int64
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
