export interface StarResult {
  fullStars: number;
  hasHalfStar: boolean;
  emptyStars: number;
}

export function computeStars(rating: number): StarResult {
  // let fullStars = Math.ceil(rating);
  // let hasHalfStar = false;
  // const emptyStars = 5 - fullStars;
  //
  // if (fullStars - rating >= 0.5) {
  //   fullStars -= 1;
  //   hasHalfStar = true;
  // }

  const fullStars = 0;
  const hasHalfStar = false;
  const emptyStars = Math.ceil(rating);

  return { fullStars, hasHalfStar, emptyStars };
}
