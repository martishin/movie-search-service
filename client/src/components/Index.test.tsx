import { expect } from "@jest/globals";
import { cleanup, render } from "@testing-library/react";

import Index from "./Index";

afterEach(() => {
  cleanup();
});
describe("Index", () => {
  test("should display Hello world!", () => {
    const { container, getByText } = render(<Index />);
    expect(container).toMatchSnapshot();

    expect(getByText("Hello world!")).toBeTruthy();
  });
});
