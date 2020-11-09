import { connect } from 'react-redux'
import Pure from './Pure'

const s = (state) => ({
  loading: state.common.loading
})

export default connect(s)(Pure)