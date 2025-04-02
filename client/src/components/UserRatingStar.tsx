import { FaRegStar, FaStar, FaStarHalfStroke } from "react-icons/fa6";

interface UserRatingStarProps {
  rating: number;
}

export default function UserRatingStar({ rating }: UserRatingStarProps) {
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
  const emptyStars = rating;

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
