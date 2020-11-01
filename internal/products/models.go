package products

import "time"

// Product representation in app
type Product struct {
	ID        int
	Name      string
	Details   string
	Active    bool
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductDiscount struct {
	ID             int
	ProductID      int
	MinQuantity    int
	MaxQuantity    int
	PriceType      PriceType
	ComboPackageID int
	Discount       int
	Active         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Stock representation in app
type Stock struct {
	ID        int
	ProductID int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Order representation in app
type Order struct {
	ID        int
	Name      string
	Details   string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ComboPackage struct {
	ID           int
	Name         string
	PackagedWith []ComboPackagedWith
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ComboPackagedWith struct {
	ID                    int
	ProductID             int
	PackagedWithProductID int
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
