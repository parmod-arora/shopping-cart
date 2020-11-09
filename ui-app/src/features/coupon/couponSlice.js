import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  coupon: {
    discount_id: '',
    expire: '',
    id: '',
    name: ''
  }
}

const couponSlice = createSlice({
  name: 'coupon',
  initialState,
  reducers: {
    populateCoupon(state, { payload }) {
      state.coupon = payload
    },
    applyCoupon() {
    },
    removeCoupon() {
    },
  }
})

export const { applyCoupon, removeCoupon, populateCoupon } = couponSlice.actions

export default couponSlice.reducer