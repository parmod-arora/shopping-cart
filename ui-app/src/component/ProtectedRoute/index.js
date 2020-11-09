import React from 'react'
import { Route, Redirect } from 'react-router-dom'
import { AUTH_TOKEN_KEY } from "../../config";

function isAuthenticated() {
  if (sessionStorage.getItem(AUTH_TOKEN_KEY)) {
    return true
  }
  return false
}

export const ProtectedRoute = ({ auth, scopes = [ ], component: Component, ...rest }) => {
  return (
    <Route {...rest} render={(props) => isAuthenticated() ? (<Component
      {...rest} />
    ) : (
      <Redirect to={{ pathname: '/login' }} />
    )
    } />
  )
}
