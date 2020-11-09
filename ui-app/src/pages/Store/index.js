import { compose } from "redux";
import { connect } from 'react-redux'
import { loadProducts, addToCart } from "../../features/products/productsSlice";
import Products from "./Store";

const s = (state) => ({
  products: state.products.products
})
const d = (dispatch) => ({
  loadProducts: compose(dispatch, loadProducts),
  addProductToCart: compose(dispatch, addToCart)
})

export default connect(s,d)(Products)