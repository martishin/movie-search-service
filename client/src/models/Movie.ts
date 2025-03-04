import Genre from "./Genre";

export default interface Movie {
  id: number;
  title: string;
  release_date: string;
  runtime: string;
  mpaa_rating: string;
  description: string;
  image: string;
  video: string;
  genres: Genre[];
  user_rating: number;
  is_liked: boolean;
}
