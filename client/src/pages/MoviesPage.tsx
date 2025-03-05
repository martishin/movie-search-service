import { useEffect, useState } from "react";
import { FaSortAmountDown, FaSortAmountUp } from "react-icons/fa";
import { FaHeart, FaRegHeart } from "react-icons/fa6";
import { Link } from "react-router";

import { API_URL } from "../api";
import GenreTag from "../components/GenreTag";
import UserRatingStar from "../components/UserRatingStar";
import { useAlert } from "../context/AlertContext";
import { useAuth } from "../context/AuthContext";
import Genre from "../models/Genre";
import Movie from "../models/Movie";

export default function MoviesPage() {
  const [movies, setMovies] = useState<Movie[]>([]);
  const [filteredMovies, setFilteredMovies] = useState<Movie[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const { showAlert } = useAlert();
  const { userDetails } = useAuth();
  const [searchQuery, setSearchQuery] = useState("");
  const [sortField, setSortField] = useState<"title" | "releaseDate" | "userRating">("userRating");
  const [sortDirection, setSortDirection] = useState<"asc" | "desc">("desc");
  const [fetchError, setFetchError] = useState<string | null>(null);

  useEffect(() => {
    const fetchMovies = async () => {
      if (fetchError) return;
      setIsLoading(true);

      const apiEndpoint = userDetails
        ? `${API_URL}/api/movies-with-likes`
        : `${API_URL}/api/movies`;

      try {
        const response = await fetch(apiEndpoint, { credentials: "include" });
        if (!response.ok) throw new Error("Failed to fetch movies");

        const data = await response.json();

        const movies = data.map(
          (movie: any) =>
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

        setMovies(movies);
        setFilteredMovies(movies);
      } catch (err: unknown) {
        const errorMsg = err instanceof Error ? err.message : "An unknown error occurred";
        setFetchError(errorMsg);
        showAlert(errorMsg);
      } finally {
        setIsLoading(false);
      }
    };

    fetchMovies();
  }, [userDetails]);

  const handleSort = (field: "title" | "releaseDate" | "userRating") => {
    if (field === sortField) {
      setSortDirection(sortDirection === "asc" ? "desc" : "asc");
    } else {
      setSortField(field);
      setSortDirection("desc");
    }
  };

  const toggleLike = async (movieID: number, isLiked: boolean) => {
    if (!userDetails) return;

    const method = isLiked ? "DELETE" : "POST";
    const endpoint = `${API_URL}/api/movies/likes/${movieID}`;

    setMovies((prevMovies) =>
      prevMovies.map((movie) =>
        movie.id === movieID
          ? new Movie(
              movie.id,
              movie.title,
              movie.releaseDate,
              movie.runtime,
              movie.mpaaRating,
              movie.description,
              movie.image,
              movie.video,
              movie.genres,
              movie.userRating,
              !isLiked,
            )
          : movie,
      ),
    );

    try {
      const response = await fetch(endpoint, {
        method,
        credentials: "include",
        headers: { "Content-Type": "application/json" },
      });

      if (!response.ok) throw new Error(`Failed to ${isLiked ? "unlike" : "like"} movie`);

      setFilteredMovies((prevMovies) =>
        prevMovies.map((movie) =>
          movie.id === movieID
            ? new Movie(
                movie.id,
                movie.title,
                movie.releaseDate,
                movie.runtime,
                movie.mpaaRating,
                movie.description,
                movie.image,
                movie.video,
                movie.genres,
                movie.userRating,
                !isLiked,
              )
            : movie,
        ),
      );
    } catch (err: unknown) {
      showAlert(err instanceof Error ? err.message : "An unknown error occurred");

      setMovies((prevMovies) =>
        prevMovies.map((movie) =>
          movie.id === movieID
            ? new Movie(
                movie.id,
                movie.title,
                movie.releaseDate,
                movie.runtime,
                movie.mpaaRating,
                movie.description,
                movie.image,
                movie.video,
                movie.genres,
                movie.userRating,
                isLiked,
              )
            : movie,
        ),
      );
    }
  };

  const sortedMovies = [...filteredMovies].sort((a, b) => {
    if (sortField === "releaseDate") {
      return sortDirection === "asc"
        ? a.releaseDate.getTime() - b.releaseDate.getTime()
        : b.releaseDate.getTime() - a.releaseDate.getTime();
    }

    if (sortField === "userRating") {
      return sortDirection === "asc"
        ? a.userRating - b.userRating || a.id - b.id
        : b.userRating - a.userRating || a.id - b.id;
    }

    return (
      a[sortField].toLowerCase().localeCompare(b[sortField].toLowerCase()) *
      (sortDirection === "asc" ? 1 : -1)
    );
  });

  const handleSearch = (query: string) => {
    setSearchQuery(query);
    const lowerQuery = query.toLowerCase();
    const filtered = movies.filter(
      (movie) =>
        movie.title.toLowerCase().includes(lowerQuery) ||
        movie.genres.some((genre) => genre.genre.toLowerCase().includes(lowerQuery)),
    );
    setFilteredMovies(filtered);
  };

  const renderSortIcon = (field: "title" | "releaseDate" | "userRating") => (
    <span className="inline-block w-4">
      {sortField === field && (sortDirection === "asc" ? <FaSortAmountUp /> : <FaSortAmountDown />)}
    </span>
  );

  const renderGenres = (genres: Genre[]) => (
    <div className="flex flex-wrap gap-1">
      {genres.map((genre) => (
        <GenreTag key={genre.id} genre={genre} />
      ))}
    </div>
  );

  return (
    <div className="px-6 sm:px-8 lg:px-10">
      {/* Page Header & Search */}
      <div className="sm:flex sm:flex-col sm:items-start">
        <div className="sm:flex-auto">
          <h1 className="text-xl font-semibold text-gray-900">Movies</h1>
          <p className="mt-2 text-sm text-gray-600">
            Browse a collection of movies and view details.
          </p>
        </div>
        <div className="relative mt-4 w-64">
          <input
            type="text"
            placeholder="Search by title or genre..."
            className="block w-full rounded-md border-0 px-3 py-1.5 text-sm text-gray-900 ring-1 ring-gray-300 ring-inset placeholder:text-gray-400 focus:ring-gray-300 focus:outline-none"
            value={searchQuery}
            onChange={(e) => handleSearch(e.target.value)}
          />
          {searchQuery && (
            <button
              onClick={() => handleSearch("")}
              className="absolute inset-y-0 right-2 flex items-center text-gray-500 hover:text-gray-700"
              aria-label="Clear search"
            >
              ✕
            </button>
          )}
        </div>
      </div>

      {isLoading ? null : sortedMovies.length === 0 ? ( // <p className="mt-6 text-center text-gray-500">Loading movies...</p>
        <p className="mt-6 text-center text-gray-500">No movies found.</p>
      ) : (
        <div className="overflow-hidden">
          <table className="min-w-full text-left text-sm text-gray-700">
            <thead className="bg-transparent text-gray-900">
              <tr>
                <th scope="col" className="w-40 py-3 pr-3 pl-4 sm:pl-6">
                  Poster
                </th>
                <th
                  scope="col"
                  className="w-2/5 cursor-pointer px-3 py-3"
                  onClick={() => handleSort("title")}
                >
                  <div className="flex items-center gap-2">Movie {renderSortIcon("title")}</div>
                </th>
                {userDetails ? (
                  <th scope="col" className="w-1/12 px-3 py-3">
                    Like
                  </th>
                ) : null}
                <th
                  scope="col"
                  className="w-1/6 cursor-pointer px-3 py-3"
                  onClick={() => handleSort("userRating")}
                >
                  <div className="flex items-center gap-2">
                    User Rating {renderSortIcon("userRating")}
                  </div>
                </th>
                <th scope="col" className="w-1/6 px-3 py-3">
                  Genres
                </th>
                <th
                  scope="col"
                  className="w-1/5 cursor-pointer px-3 py-3"
                  onClick={() => handleSort("releaseDate")}
                >
                  <div className="flex items-center gap-2">
                    Release Date {renderSortIcon("releaseDate")}
                  </div>
                </th>
                <th scope="col" className="w-1/12 px-3 py-3">
                  MPA Rating
                </th>
              </tr>
            </thead>
            <tbody>
              {sortedMovies.map((movie) => (
                <tr key={movie.id} className="border-t border-gray-200">
                  <td className="px-3 py-3 whitespace-nowrap">
                    <Link to={`/movies/${movie.id}`}>
                      <img
                        src={movie.formattedImage}
                        alt={movie.title}
                        className="h-30 min-w-20 object-cover transition-transform hover:scale-105"
                      />
                    </Link>
                  </td>
                  <td className="px-3 py-3 whitespace-nowrap">
                    <Link to={`/movies/${movie.id}`} className="text-blue-600 hover:underline">
                      {movie.title}
                    </Link>
                  </td>
                  {userDetails && (
                    <td className="px-3 py-3">
                      <button onClick={() => toggleLike(movie.id, movie.isLiked)}>
                        {movie.isLiked ? <FaHeart className="text-red-500" /> : <FaRegHeart />}
                      </button>
                    </td>
                  )}
                  <td className="px-3 py-3 whitespace-nowrap">
                    <UserRatingStar rating={movie.userRating} />
                  </td>
                  <td className="px-3 py-3 whitespace-nowrap">{renderGenres(movie.genres)}</td>
                  <td className="px-3 py-3 whitespace-nowrap">{movie.formattedReleaseDate}</td>
                  <td className="px-3 py-3 whitespace-nowrap">{movie.mpaaRating}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
