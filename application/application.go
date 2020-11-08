package application

import (
	"cinemo.com/shoping-cart/internal/cart"
	"cinemo.com/shoping-cart/internal/coupons"
	"cinemo.com/shoping-cart/internal/discounts"
	"cinemo.com/shoping-cart/internal/products"
	"cinemo.com/shoping-cart/internal/users"
)

// App represents shopping cart application
type App struct {
	UserService     users.Service
	ProductService  products.Service
	DiscountService discounts.Service
	CartService     cart.Service
	CouponService   coupons.Service
}
