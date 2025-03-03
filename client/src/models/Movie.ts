export default interface Movie {
  id: number;
  title: string;
  release_date: string;
  runtime: string;
  mpaa_rating: string;
  description: string;
  image: string;
  genres: { id: number; genre: string }[];
  user_rating: number;
}
