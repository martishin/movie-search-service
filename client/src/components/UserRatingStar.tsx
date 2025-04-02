import { FaRegStar, FaStar, FaStarHalfStroke } from "react-icons/fa6";

import { computeStars } from "./utils/computeStars";

interface UserRatingStarProps {
  rating: number;
}

export default function UserRatingStar({ rating }: UserRatingStarProps) {
  const { fullStars, hasHalfStar, emptyStars } = computeStars(rating);

  return (
    <div className="flex">
      {Array.from({ length: fullStars }).map((_, i) => (
        <FaStar key={`full-${i}`} className="text-yellow-500" />
      ))}
      {hasHalfStar && <FaStarHalfStroke key="half" className="text-yellow-500" />}
      {Array.from({ length: emptyStars }).map((_, i) => (
        <FaRegStar key={`empty-${i}`} className="text-yellow-500" />
      ))}
    </div>
  );
}
