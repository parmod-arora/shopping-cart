import { createStore, applyMiddleware, compose } from '@reduxjs/toolkit'
import createSagaMiddleware from 'redux-saga'
import { rootSaga, rootReducer } from "../features/root";
import { createBrowserHistory } from 'history'
import { routerMiddleware } from 'react-router-redux'

const sagaMiddleware = createSagaMiddleware()

export const history = createBrowserHistory()
const reduxRouterMiddleware = routerMiddleware(history)

let composeEnhancers = compose
if (process.env.NODE_ENV !== 'production') {
  composeEnhancers = typeof window === 'object' && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ ? window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__({}) : compose
}

export default function configureStore() {
  const store = createStore(
    rootReducer,
    composeEnhancers(
      applyMiddleware(
        sagaMiddleware,
        reduxRouterMiddleware
      )
    )
  );

  sagaMiddleware.run(rootSaga)
  return store;
}