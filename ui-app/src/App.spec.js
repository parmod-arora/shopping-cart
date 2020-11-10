import React from "react";
import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";
import { Provider } from 'react-redux';
import { ThemeProvider } from '@material-ui/core/styles';
import configureStore from "./store/store";
import theme from './theme';
import App from "./App";

const store = configureStore()

let container = null;
beforeEach(() => {
  // setup a DOM element as a render target
  container = document.createElement("div");
  document.body.appendChild(container);
});

afterEach(() => {
  // cleanup on exiting
  unmountComponentAtNode(container);
  container.remove();
  container = null;
});

function TestComp() {
  return (<ThemeProvider theme={theme}>
    <Provider store={store}>
      <App />
    </Provider>
  </ThemeProvider>)
}

it("renders app", () => {
  act(() => {
    render(<TestComp />, container);
  });
  expect(container).toBeTruthy()
});