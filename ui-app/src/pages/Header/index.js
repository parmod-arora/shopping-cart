import { connect } from 'react-redux'
import { compose } from "redux";
import { Header } from './Header'
import { logout } from "../../features/auth/loginSlice";

const s = (state) => ({
  cartItemsCount: (state.products.cartItems &&  state.products.cartItems.length > 0) ?  state.products.cartItems.reduce((count,item) => {
    return count + item.quantity
  }, 0) : ''
})
const d = (dispatch) => ({
  logout: compose(dispatch, logout),
})
export default connect(s, d)(Header)