package application

import (
	"cinemo.com/shoping-cart/internal/cart"
	"cinemo.com/shoping-cart/internal/discount"
	"cinemo.com/shoping-cart/internal/products"
	"cinemo.com/shoping-cart/internal/users"
)

// App represents shopping cart application
type App struct {
	UserService     users.Service
	ProductService  products.Service
	DiscountService discount.Service
	CartService     cart.Service
}
