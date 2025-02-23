import { cleanup, render } from "@testing-library/react";
import App from "./App";
import { expect } from "@jest/globals";

afterEach(() => {
  cleanup();
});
describe("App", () => {
  test("should display Hello world!", () => {
    const { container, getByText } = render(<App />);
    expect(container).toMatchSnapshot();

    expect(getByText("Hello world!")).toBeTruthy();
  });
});
