import React from 'react'
import PropTypes from 'prop-types'
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';

function Pure({ succssMsg, errorMsg }) {
  if (succssMsg) {
    return <Snackbar anchorOrigin={{ 
      vertical: 'top',
      horizontal: 'right' 
    }} open={!!succssMsg}>
      <Alert elevation={6} variant="filled" severity="success">
        {succssMsg}
      </Alert>
    </Snackbar>
  }
  if (errorMsg) {
    return <div>
        <Snackbar open={!!errorMsg}>
          <Alert elevation={6} variant="filled" severity="error">
            {errorMsg}
          </Alert>
      </Snackbar>
    </div>
  }
  return <div />
}

Pure.propTypes = {
  succssMsg: PropTypes.string,
  errorMsg: PropTypes.string
}
Pure.defaultProps = {
  succssMsg: "",
  errorMsg: ""
}

export default Pure