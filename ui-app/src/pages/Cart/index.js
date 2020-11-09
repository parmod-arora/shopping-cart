import { compose } from "redux";
import { connect } from 'react-redux'
import { getUserCart } from "../../features/products/productsSlice";
import { addToCart, removeFromCart, checkout } from "../../features/products/productsSlice";
import { applyCoupon, removeCoupon } from "../../features/coupon/couponSlice";
import Checkout from "./Checkout";

const s = (state) => {
  return {
    lineItems: state.products.lineItems,
    cartItems: state.products.cartItems,
    subTotalAmount: state.products.subTotalAmount,
    totalSavingAmount: state.products.totalSavingAmount,
    totalAmount: state.products.totalAmount,
    coupon: state.coupon.coupon
  }
}
const d = (dispatch) => ({
  fetchUserCart: compose(dispatch, getUserCart),
  addProductToCart: compose(dispatch, addToCart),
  removeProductfromCart: compose(dispatch, removeFromCart),
  dispatchApplyCoupon: compose(dispatch, applyCoupon),
  dispatchRemoveCoupon: compose(dispatch, removeCoupon),
  dispatchCheckout: compose(dispatch, checkout)
})

export default connect(s, d)(Checkout)