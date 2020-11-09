import { takeEvery, takeLatest } from 'redux-saga/effects'

export const takeEverySagaWatcher = function (sagas) {
  const result = []
  for (const k in sagas) {
    if (sagas.hasOwnProperty(k)) {
      const v = sagas[k]
      result.push(function * () {
        yield takeEvery(k, v)
      }())
    }
  }
  return result
}
// Can be useful to handle AJAX requests where we want to only have the response to the latest request.
export const takeLatestSagaWatcher = function (sagas) {
  const result = []
  for (const k in sagas) {
    if (sagas.hasOwnProperty(k)) {
      const v = sagas[k]
      result.push(function * () {
        yield takeLatest(k, v)
      }())
    }
  }
  return result
}