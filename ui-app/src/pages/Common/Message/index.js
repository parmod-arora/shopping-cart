import { connect } from 'react-redux'
import Pure from './Pure'

const s = (state) => ({
  succssMsg: state.common.succssMsg,
  errorMsg: state.common.errorMsg
})

export default connect(s)(Pure)