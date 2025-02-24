import { cleanup, render } from "@testing-library/react";
import Index from "./Index";
import { expect } from "@jest/globals";

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
