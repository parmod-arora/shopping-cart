import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  succssMsg: '',
  errorMsg: '',
  loading: 0
}

const commonSlice = createSlice({
  name: 'common',
  initialState,
  reducers: {
    successMsg(state, action) {
      state.succssMsg = action.payload
    },
    errorMsg(state, action) {
      state.errorMsg = action.payload
    },
    clearSuccesMsg(state) {
      state.succssMsg = initialState.succssMsg
    },
    clearErrorMsg(state) {
      state.errorMsg = initialState.errorMsg
    },
    startLoading(state){
      state.loading = state.loading +1
    },
    stopLoading(state) {
      state.loading = (state.loading - 1) < 0 ? 0 : state.loading - 1
    }
  }
})

export const { successMsg, errorMsg, clearSuccesMsg, clearErrorMsg, startLoading, stopLoading } = commonSlice.actions

export default commonSlice.reducer