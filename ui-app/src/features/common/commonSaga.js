import { put, delay } from 'redux-saga/effects'
import { errorMsg, successMsg, clearSuccesMsg, clearErrorMsg } from "../common/commonSlice";
import { takeEverySagaWatcher } from "../../store/framework";

export const sagas = {
  [successMsg]: function * () {
    yield delay(3000)
    yield put(clearSuccesMsg())
  },
  [errorMsg]: function * () {
    yield delay(3000)
    yield put(clearErrorMsg())
  }
}

export const commonSagaWatcher = takeEverySagaWatcher(sagas)