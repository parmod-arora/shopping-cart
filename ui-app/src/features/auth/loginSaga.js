import { put, call } from 'redux-saga/effects'
import { push } from 'react-router-redux'
import { startLoading, stopLoading, errorMsg } from "../common/commonSlice";
import axios from "../../axios-instance";
import { login, loginSuccess, logout } from './loginSlice';
import { API_URL, AUTH_TOKEN_KEY } from "../../config";
import { takeEverySagaWatcher } from "../../store/framework";

const login_url = `${API_URL}/api/v1/users/login`
export const sagas = {
  [loginSuccess]: function * ({ payload }) {
    sessionStorage.setItem(AUTH_TOKEN_KEY, payload.token)
    yield put(push('/products'))
  },
  [login]: function * ({ payload }) {
    try {
      yield put(startLoading())
      const response = yield call(axios.post, login_url, payload)
      yield put(loginSuccess(response))
    } catch (error) {
      if (error.error_description) {
        yield put(errorMsg(error.error_description))
      } else {
        yield put(errorMsg('Unexpected error from server!'))  
      }
    }
    yield put(stopLoading())
  },
  [logout]: function * ({ payload }) {
    sessionStorage.removeItem(AUTH_TOKEN_KEY)
    yield put(push('/login'))
  }
}

export const loginSagaWatcher = takeEverySagaWatcher(sagas)
