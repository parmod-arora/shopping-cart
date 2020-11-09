import { compose } from "redux";
import { connect } from 'react-redux'
import Form from "./Form";
import { login } from "../../features/auth/loginSlice";

const s = (state) => ({
  state: state.login
})
const d = (dispatch) => ({
  login: compose(dispatch, login)
})

export default connect(s,d)(Form)