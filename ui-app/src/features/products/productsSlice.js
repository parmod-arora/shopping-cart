import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  cartId: '',
  products:[],
  cartItems: [],
  lineItems: [],
  subTotalAmount: '',
  totalSavingAmount: '',
  totalAmount: ''
}

const productsSlice = createSlice({
  name: 'products',
  initialState,
  reducers: {
    loadProducts(){},
    loadProductsSuccess(state, action) {
      state.products = action.payload
    },
    loadProductsFailure(state, action) {
      state.products = initialState
    },
    addToCart() {
    },
    removeFromCart() {
    },
    checkout(){
    },
    populateCartInfo(state, { payload }) {
      state.cartId =  payload["id"]
      const cartItems =  payload["cart_items"] || []
      state.cartItems = cartItems.sort(function(a, b){
        return a.product.id - b.product.id
      })
      const lineItems = payload["line_items"] || []
      state.lineItems = lineItems.sort(function(a, b){
        return a.discount_applied.id - b.discount_applied.id
      })
      state.subTotalAmount = payload["sub_total_amount"] || null
      state.totalSavingAmount = payload["total_saving_amount"] || null
      state.totalAmount = payload["total_amount"] || null
    },
    getUserCart(){
    }
  }
})

export const { 
  loadProducts,
  loadProductsSuccess,
  loadProductsFailure, 
  addToCart,
  addToCartSuccess,
  populateCartInfo,
  getUserCart,
  removeFromCart,
  checkout
} = productsSlice.actions
export default productsSlice.reducer