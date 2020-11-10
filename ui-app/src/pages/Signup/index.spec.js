import React from "react";
import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";
import Signup from "./Form";

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
  return <Signup />
}

it("should renders signup form", () => {
  act(() => {
    render(<TestComp />, container);
  });
  expect(container).toBeTruthy()
});
