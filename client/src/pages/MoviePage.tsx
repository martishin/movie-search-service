import { ReactNode, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router";
import { FaArrowLeft } from "react-icons/fa6";

import Movie from "../models/Movie";
import GenreTag from "../components/GenreTag";
import UserRatingStar from "../components/UserRatingStar";

export default function MoviePage(): ReactNode {
  const [movie, setMovie] = useState<Movie | null>(null);
  const navigate = useNavigate();
  const { id } = useParams();
  const [isFetchingMovie, setIsFetchingMovie] = useState(true);

  useEffect(() => {
    const fetchMovie = async () => {
      try {
        const response = await fetch(`/api/movies/${id}`, {
          method: "GET",
          headers: { "Content-Type": "application/json" },
        });
        if (!response.ok) throw new Error("Failed to fetch movie");

        const data: Movie = await response.json();
        const formattedMovie = {
          ...data,
          release_date: new Date(data.release_date).toLocaleDateString("en-US", {
            year: "numeric",
            month: "long",
            day: "numeric",
          }),
        };
        setMovie(formattedMovie);
      } catch (error) {
        console.error("Error fetching movie:", error);
      } finally {
        setIsFetchingMovie(false);
      }
    };

    fetchMovie();
  }, [id]);

  return (
    <div className="px-6 sm:px-8 lg:px-10">
      <button
        onClick={() => navigate("/movies")}
        className="mb-4 flex items-center gap-2 text-blue-600 hover:text-blue-800"
      >
        <FaArrowLeft className="h-5 w-5" /> Back
      </button>
      {isFetchingMovie ? (
        <p className="mt-4 text-center text-gray-500">Loading movie...</p>
      ) : (
        movie && (
          <div className="mt-6 flex flex-col gap-6 md:flex-row">
            {movie.image && (
              <div className="w-40 flex-shrink-0 md:w-48">
                <img
                  src={`https://image.tmdb.org/t/p/w400/${movie.image}`}
                  alt={movie.title}
                  className="h-auto w-full rounded shadow-md"
                />
              </div>
            )}
            <div className="flex-1 space-y-2 text-gray-900">
              <h3 className="text-2xl font-bold tracking-tight">{movie.title}</h3>
              <p className="flex items-center gap-2 font-semibold">
                <span className="text-gray-700">User Rating:</span>
                <UserRatingStar rating={movie.user_rating} />
              </p>
              <p className="flex items-center gap-2 font-semibold">
                <span className="text-gray-700">Genres:</span>
                <span className="flex flex-wrap gap-2">
                  {movie.genres.map((g) => (
                    <GenreTag key={g.id} genre={g} />
                  ))}
                </span>
              </p>
              <p className="font-semibold">
                <span className="text-gray-700">Release Date:</span>
                <span className="font-normal"> {movie.release_date}</span>
              </p>
              <p className="font-semibold">
                <span className="text-gray-700">Length:</span>
                <span className="font-normal"> {movie.runtime} minutes</span>
              </p>
              <p className="font-semibold">
                <span className="text-gray-700">MPAA Rating:</span>
                <span className="font-normal"> {movie.mpaa_rating}</span>
              </p>
            </div>
          </div>
        )
      )}
      {movie && (
        <div className="mt-6">
          <h4 className="text-lg font-semibold text-gray-900">Description</h4>
          <p className="leading-relaxed text-gray-800">{movie.description}</p>
        </div>
      )}
    </div>
  );
}
