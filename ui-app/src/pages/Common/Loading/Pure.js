import React from 'react'
import PropTypes from 'prop-types'
import CircularProgress from '@material-ui/core/CircularProgress';
import { makeStyles } from '@material-ui/core/styles';

function Pure({ loading }) {
  const classes = makeStyles((theme) => ({
    wrapper: {
        zIndex: 1000,
        display: "flex",
        position: "fixed",
        top: 0,
        right: 0,
        bottom: 0,
        left: 0,
        width: "100%",
        height: "100%"
    },
    marginAutoItem: {
      margin: 'auto'
    },
  }))()

  return (
    loading ? <div id="loader" className={classes.wrapper}><CircularProgress className={classes.marginAutoItem} /></div> : <div />
  )
}

Pure.propTypes = {
  loading: PropTypes.number
}
Pure.defaultProps = {
  loading: 0
}

export default Pure