import { describe, expect, test } from "@jest/globals";

import { sum } from "./utils";

describe("Sum function", () => {
  test("Returns correct value", () => {
    expect(sum(2, 3)).toEqual(5);
  });
});
