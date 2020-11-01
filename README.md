# shopping-cart


## Task
Develop an online eCommerce store selling fruit, which contains the following features:

1. Simple sign-up and login form.
1. Browse the following products
  - Apples
  - Bananas
  - Pears
  - Oranges
1. Add items to your cart
  - Adjust quantity.
  - Delete items from the cart.
  - Apply coupons.
Checkout your cart.
Mocked purchase (a payment gateway is not required, but a route must exist in the backend validating the payment).
An address does not need to be entered.
Requirements

 Architecture diagrams.

Single-page frontend app (cannot use an existing online stores such as Prestashop).
Backend RESTful web service written in GoLang.
Users must be able to return to their cart after closing the browser, and see the previous items that were added.


### Cart Rules

1. If 7 or more apples are added to the cart, a 10% discount is applied to all apples.

1. For each set of 4 pears and 2 bananas, a 30% discount is applied, to each set. (These sets must be added to their own cart item entry.) If pears or bananas already exist in the cart, this discount must be recalculated when new pears or bananas are added.

1. A coupon code can be used to get a 30% discount on oranges, if applied to the cart, otherwise oranges are full price.
  - Can only be applied once.
  - Has an configurable expiry timeout (10 seconds for testing purposes) once generated.
1.  The following totals must be shown:
  - Total price.
  - Total savings.