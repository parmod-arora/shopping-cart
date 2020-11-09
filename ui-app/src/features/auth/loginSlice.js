import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  token: "",
  loading: false,
  error: []
}

const loginSlice = createSlice({
  name: 'login',
  initialState,
  reducers: {
    login(state) {
      state.error = []
      state.loading = true
    },
    loginSuccess(state, action) {
      state.loading = false
    },
    loginFailure(state, action) {
      state.loading = false
    },
    logout(){}
  }
})

export const { login, loginSuccess, loginFailure, logout } = loginSlice.actions

export default loginSlice.reducer