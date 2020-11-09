import { combineReducers } from 'redux';
import { all } from 'redux-saga/effects'
import { routerReducer} from 'react-router-redux'
// reducers
import products  from "./products/productsSlice";
import login from "./auth/loginSlice";
import signup from "./auth/signupSlice";
import common from "./common/commonSlice";
import coupon from "./coupon/couponSlice";
// sagas
import { signupSagaWatcher } from "./auth/signupSaga";
import { loginSagaWatcher } from "./auth/loginSaga";
import { commonSagaWatcher } from "./common/commonSaga";
import { productsSagaWatcher } from "./products/productsSaga";
import { couponSagaWatcher } from "./coupon/couponSaga";

export function* rootSaga () {
  yield all([
    ...signupSagaWatcher,
    ...commonSagaWatcher,
    ...loginSagaWatcher,
    ...productsSagaWatcher,
    ...couponSagaWatcher
  ])
}

export const rootReducer = combineReducers({
  common,
  router: routerReducer,
  products,
  login,
  signup,
  coupon
});