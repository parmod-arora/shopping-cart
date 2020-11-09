import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  values: {
    email: '',
    password: '',
    firstName: '',
    lastName: '',
  },
  loading: false
}

const signupSlice = createSlice({
  name: 'signup',
  initialState,
  reducers: {
    signup(){},
    signupSuccess(state, action) {
      console.log(action)
      state.loading = false
    },
    signupFailure(state, action) {
      console.log(action)
      state.loading = false
    }
  }
})

export const { signup, signupSuccess, signupFailure } = signupSlice.actions

export default signupSlice.reducer