import { put, call } from 'redux-saga/effects'
import { push } from 'react-router-redux'
import { startLoading, stopLoading, errorMsg, successMsg } from "../common/commonSlice";
import axios from "../../axios-instance";
import { signup } from './signupSlice';
import { API_URL } from "../../config";
import { takeEverySagaWatcher } from "../../store/framework";
const signup_url = `${API_URL}/api/v1/users/signup`

export const sagas = {
  [signup]: function * ({ payload }) {
    try {
      yield put(startLoading())
      yield call(axios.post, signup_url, payload)
      yield put(successMsg('Thank you for signing up!'))
      yield put(push('/'))
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

export const signupSagaWatcher = takeEverySagaWatcher(sagas)