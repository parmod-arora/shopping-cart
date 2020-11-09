import { compose } from "redux";
import { connect } from 'react-redux'
import Form from "./Form";
import { signup } from "../../features/auth/signupSlice";

const s = (state) => ({
  state: state.signup,
  notificationState: state.notification
})
const d = (dispatch) => ({
  signup: compose(dispatch, signup)
})

export default connect(s,d)(Form)
