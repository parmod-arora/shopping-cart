import { put, call, select } from "redux-saga/effects";
import { takeEverySagaWatcher } from "../../store/framework";
import axios from "../../axios-instance";
import { startLoading, stopLoading, errorMsg, successMsg } from "../common/commonSlice";
import {
  loadProducts,
  loadProductsFailure,
  loadProductsSuccess,
  addToCart,
  removeFromCart,
  populateCartInfo,
  getUserCart,
  checkout
} from "./productsSlice";
import { populateCoupon } from "../coupon/couponSlice";
import { API_URL } from "../../config";

const products_url = `${API_URL}/api/v1/products/`
const carts_url = `${API_URL}/api/v1/carts/items`
const checkout_url = `${API_URL}/api/v1/carts/checkout`

const selectCartItems = state => state.products.cartItems
const selectUserCartId = state => state.products.cartId
const selectCartTotalAmount = state => state.products.totalAmount

export const sagas = {
  [loadProducts]: function* ({ payload }) {
    try {
      yield put(getUserCart())
      yield put(startLoading())
      const response = yield call(axios.get, products_url)
      yield put(loadProductsSuccess(response))
    } catch (error) {
      yield put(loadProductsFailure())
    }
    yield put(stopLoading())
  },
  [loadProductsFailure]: function* ({ payload }) {
    yield put(errorMsg('Unexpected error from server!'))
  },
  [addToCart]: function* ({ payload }) {
    const cartItems = yield select(selectCartItems)
    const request = prepareRequest(payload, cartItems, 1)
    try {
      yield put(startLoading())
      const response = yield call(axios.put, carts_url, request)
      yield put(populateCartInfo(response))
    } catch (error) {
      yield put(errorMsg('Unexpected error from server!'))
    }
    yield put(stopLoading())
  },
  [removeFromCart]: function* ({ payload }) {
    const cartItems = yield select(selectCartItems)
    const request = prepareRequest(payload, cartItems, -1)
    try {
      yield put(startLoading())
      const response = yield call(axios.put, carts_url, request)
      yield put(populateCartInfo(response))
    } catch (error) {
      yield put(errorMsg('Unexpected error from server!'))
    }
    yield put(stopLoading())
  },
  [getUserCart]: function* () {
    try {
      yield put(startLoading())
      const response = yield call(axios.get, carts_url)
      yield put(populateCartInfo(response))
    } catch (error) {
      yield put(errorMsg('Unexpected error from server!'))
    }
    yield put(stopLoading())
  },
  [populateCartInfo]: function * ({payload}) {
    yield put(populateCoupon(payload.coupon)) 
  },
  [checkout]: function* ({ payload }) {
    try {
      const cartId = yield select(selectUserCartId)
      const amount = yield select(selectCartTotalAmount)
      yield put(startLoading())
      const request = {
        cart_id: cartId,
        amount: amount
      }
      yield call(axios.post, checkout_url, request)
      yield put(getUserCart())
      yield put(successMsg('Thank you for shopping with you'))
    } catch (error) {
      yield put(errorMsg('Unexpected error from server!'))
    }
    yield put(stopLoading())
  }
}

export const productsSagaWatcher = takeEverySagaWatcher(sagas)

function prepareRequest(product, cartItems, delta) {
  let count = delta
  if (cartItems && cartItems.length > 0) {
    for (let index = 0; index < cartItems.length; index++) {
      const cartItem = cartItems[index];
      if (cartItem.product && cartItem.product.id === product.id) {
        count = cartItem.quantity + delta
      }
    }
  }
  return {
    "product_id": product.id,
    "quantity": count
  }
}