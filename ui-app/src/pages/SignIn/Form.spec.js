import React from "react";
import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";
import Signin from "./Form";

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
  return <Signin />
}

it("should renders signin form", () => {
  act(() => {
    render(<TestComp />, container);
  });
  expect(container).toBeTruthy()
});
