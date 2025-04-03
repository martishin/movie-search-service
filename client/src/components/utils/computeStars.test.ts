import { computeStars } from "./computeStars";

describe("computeStars", () => {
  it("should return 5 full stars for rating 5", () => {
    expect(computeStars(5)).toEqual({ fullStars: 5, hasHalfStar: false, emptyStars: 0 });
  });

  it("should return 4 full and 1 half star for rating 4.5", () => {
    expect(computeStars(4.5)).toEqual({ fullStars: 4, hasHalfStar: true, emptyStars: 0 });
  });

  it("should round up to 5 full stars for rating 4.8", () => {
    expect(computeStars(4.8)).toEqual({ fullStars: 5, hasHalfStar: false, emptyStars: 0 });
  });

  it("should return 3 full, 1 half, 1 empty for rating 3.5", () => {
    expect(computeStars(3.5)).toEqual({ fullStars: 3, hasHalfStar: true, emptyStars: 1 });
  });

  it("should return 0 full, 0 half, 5 empty for rating 0", () => {
    expect(computeStars(0)).toEqual({ fullStars: 0, hasHalfStar: false, emptyStars: 5 });
  });

  it("should return 2 full, 1 half, 2 empty for rating 2.1", () => {
    expect(computeStars(2.1)).toEqual({ fullStars: 2, hasHalfStar: true, emptyStars: 2 });
  });

  it("should return 3 full, 0 half, 2 empty for rating 2.6", () => {
    expect(computeStars(2.6)).toEqual({ fullStars: 3, hasHalfStar: false, emptyStars: 2 });
  });
});
