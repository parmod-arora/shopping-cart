import "babel-polyfill"
import React from "react";
import { ThemeProvider } from '@material-ui/core/styles';
import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";
import Checkout from "./Checkout";
import { Provider } from 'react-redux';
import configureStore from "../../store/store";
import theme from '../../theme';
import { MemoryRouter } from "react-router-dom";

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
      <MemoryRouter>
        <Checkout />
      </MemoryRouter>
    </Provider>
  </ThemeProvider>)
}

it("should renders Checkout comp", () => {
  act(() => {
    render( <TestComp />, container);
  });
  expect(container).toBeTruthy()
});
