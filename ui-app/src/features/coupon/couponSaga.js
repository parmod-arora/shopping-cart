import { put, call, select } from 'redux-saga/effects'
import { applyCoupon, removeCoupon } from "./couponSlice";
import { populateCartInfo } from "../products/productsSlice";
import { startLoading, stopLoading, errorMsg } from "../common/commonSlice";
import { takeEverySagaWatcher } from "../../store/framework";
import { API_URL } from "../../config";
import axios from "../../axios-instance";

const selectUserCartId = state => state.products.cartId
const add_coupon_url = `${API_URL}/api/v1/carts/coupon/add`
const remove_coupon_url = `${API_URL}/api/v1/carts/coupon/remove`
export const sagas = {
  [applyCoupon]: function * ({payload}) {
    const cartId = yield select(selectUserCartId)
    const request = {
      coupon: payload.coupon,
      cart_id: cartId
    }
    try {
      yield put(startLoading())
      const response = yield call(axios.post, add_coupon_url, request)
      yield put(populateCartInfo(response))
    } catch (error) {
      if (error.error_description) {
        yield put(errorMsg(error.error_description))
      } else {
        yield put(errorMsg('Unexpected error from server!'))  
      }
    }
    yield put(stopLoading())
  },
  [removeCoupon]: function * ({payload}) {
    const cartId = yield select(selectUserCartId)
    const request = {
      coupon_id: payload.id,
      cart_id: cartId
    }
    try {
      yield put(startLoading())
      const response = yield call(axios.post, remove_coupon_url, request)
      yield put(populateCartInfo(response))
    } catch (error) {
      if (error.error_description) {
        yield put(errorMsg(error.error_description))
      } else {
        yield put(errorMsg('Unexpected error from server!'))  
      }
    }
    yield put(stopLoading())
  }
}

export const couponSagaWatcher = takeEverySagaWatcher(sagas)