import { useEffect, useState } from "react";
import { FaPlay } from "react-icons/fa";
import { Link } from "react-router";

import { API_URL } from "../api";
import GenreTag from "../components/GenreTag";
import UserRatingStar from "../components/UserRatingStar";
import { useAlert } from "../context/AlertContext";
import Movie from "../models/Movie";

export default function WatchOnlinePage() {
  const [movies, setMovies] = useState<Movie[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const { showAlert } = useAlert();
  const [fetchError, setFetchError] = useState<string | null>(null);

  useEffect(() => {
    if (fetchError) return;
    fetch(`${API_URL}/api/public/movies`)
      .then((res) => {
        if (!res.ok) throw new Error("Failed to fetch movies");
        return res.json();
      })
      .then((data: any[]) => {
        const movies = data.map(
          (movie) =>
            new Movie(
              movie.id,
              movie.title,
              movie.release_date,
              movie.runtime,
              movie.mpaa_rating,
              movie.description,
              movie.image,
              movie.video,
              movie.genres ?? [],
              movie.user_rating ?? 0,
              movie.is_liked ?? false,
            ),
        );

        // Sort by userRating (highest first)
        const sortedMovies = [...movies].sort((a, b) => {
          if (b.userRating !== a.userRating) {
            return b.userRating - a.userRating;
          }
          return a.id - b.id;
        });
        setMovies(sortedMovies);
      })
      .catch((err) => {
        const errorMsg = err instanceof Error ? err.message : "An unknown error occurred";
        setFetchError(errorMsg);
        showAlert(errorMsg);
      })
      .finally(() => setIsLoading(false));
  }, [showAlert]);

  return (
    <div className="px-6 sm:px-8 lg:px-10">
      <h1 className="text-xl font-semibold text-gray-900">Watch Movies Online</h1>
      <p className="mt-2 text-sm text-gray-600">Select a movie to watch.</p>
      {isLoading ? null : movies.length === 0 ? ( // <p className="mt-6 text-center text-gray-500">Loading movies...</p>
        <p className="mt-6 text-center text-gray-500">No movies found.</p>
      ) : (
        <div className="mt-6 grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-5 lg:grid-cols-5 xl:grid-cols-5 2xl:grid-cols-5">
          {movies.map((movie) => (
            <Link
              key={movie.id}
              to={`/movies/${movie.id}`}
              className="group relative overflow-hidden rounded-lg shadow-md transition-transform hover:scale-105"
            >
              <img
                src={movie.formattedImage}
                alt={movie.title}
                className="h-full w-full object-cover"
              />
              <div className="absolute inset-0 flex flex-col items-center justify-end bg-gradient-to-t from-black/80 via-black/40 to-transparent p-4 opacity-0 transition-opacity group-hover:opacity-100">
                <div className="mb-2 flex flex-wrap justify-center gap-1">
                  {movie.genres.map((genre) => (
                    <GenreTag key={genre.id} genre={genre} />
                  ))}
                </div>
                <UserRatingStar rating={movie.userRating} />
                <FaPlay className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 text-5xl text-white opacity-90" />
              </div>
            </Link>
          ))}
        </div>
      )}
    </div>
  );
}
