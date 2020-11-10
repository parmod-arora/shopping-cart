import React from "react";
import { MemoryRouter } from "react-router-dom";
import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";
import Pure from "./Pure";
// import { ThemeProvider } from '@material-ui/core/styles';
// import theme from '../../theme';

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

it("should renders message comp", () => {
  act(() => {
    render(<MemoryRouter>
      <Pure />
    </MemoryRouter>, container);
  });
  expect(container).toBeTruthy()
});
