import Genre from "./Genre";

export default class Movie {
  id: number;
  title: string;
  releaseDate: Date;
  runtime: number;
  mpaaRating: string;
  description: string;
  image: string;
  video: string;
  userRating: number;
  genres: Genre[];
  isLiked: boolean;

  constructor(
    id: number,
    title: string,
    releaseDate: string | Date,
    runtime: number,
    mpaaRating: string,
    description: string,
    image: string,
    video: string,
    genres: Genre[],
    userRating: number,
    isLiked: boolean,
  ) {
    this.id = id;
    this.title = title;
    this.releaseDate = releaseDate instanceof Date ? releaseDate : new Date(releaseDate);
    this.runtime = runtime;
    this.mpaaRating = mpaaRating;
    this.description = description;
    this.image = image;
    this.video = video;
    this.genres = genres;
    this.userRating = userRating;
    this.isLiked = isLiked;
  }

  get formattedReleaseDate(): string {
    return this.releaseDate.toLocaleDateString("en-US", {
      year: "numeric",
      month: "long",
      day: "numeric",
    });
  }

  get formattedImage(): string {
    return `https://image.tmdb.org/t/p/w400/${this.image}.jpg`;
  }

  get formattedVideo(): string {
    return `https://www.youtube.com/embed/${this.video}`;
  }
}
