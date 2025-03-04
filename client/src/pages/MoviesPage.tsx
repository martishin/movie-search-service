import { useEffect, useState } from "react";
import { Link } from "react-router";
import { useAlert } from "../context/AlertContext";
import Movie from "../models/Movie";
import { FaSortAmountDown, FaSortAmountUp } from "react-icons/fa";
import Genre from "../models/Genre";
import GenreTag from "../components/GenreTag";
import UserRatingStar from "../components/UserRatingStar";

export default function MoviesPage() {
  const [movies, setMovies] = useState<Movie[]>([]);
  const [filteredMovies, setFilteredMovies] = useState<Movie[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const { showAlert } = useAlert();
  const [searchQuery, setSearchQuery] = useState("");
  const [sortField, setSortField] = useState<
    "title" | "release_date" | "mpaa_rating" | "user_rating"
  >("user_rating");
  const [sortDirection, setSortDirection] = useState<"asc" | "desc">("desc");

  useEffect(() => {
    fetch("/api/movies")
      .then((res) => {
        if (!res.ok) throw new Error("Failed to fetch movies");
        return res.json();
      })
      .then((data: Movie[]) => {
        const formattedMovies = data.map((movie) => ({
          ...movie,
          release_date: new Date(movie.release_date).toLocaleDateString("en-US", {
            year: "numeric",
            month: "long",
            day: "numeric",
          }),
        }));
        setMovies(formattedMovies);
        setFilteredMovies(formattedMovies);
      })
      .catch((err) => {
        showAlert(err.message);
      })
      .finally(() => setIsLoading(false));
  }, [showAlert]);

  const handleSort = (field: "title" | "release_date" | "mpaa_rating" | "user_rating") => {
    if (field === sortField) {
      setSortDirection(sortDirection === "asc" ? "desc" : "asc");
    } else {
      setSortField(field);
      setSortDirection("desc");
    }
  };

  const sortedMovies = [...filteredMovies].sort((a, b) => {
    if (sortField === "release_date") {
      return sortDirection === "asc"
        ? new Date(a.release_date).getTime() - new Date(b.release_date).getTime()
        : new Date(b.release_date).getTime() - new Date(a.release_date).getTime();
    }

    if (sortField === "user_rating") {
      return sortDirection === "asc"
        ? a.user_rating - b.user_rating
        : b.user_rating - a.user_rating;
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

  const renderSortIcon = (field: "title" | "release_date" | "mpaa_rating" | "user_rating") => (
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

      {isLoading ? (
        <p className="mt-6 text-center text-gray-500">Loading movies...</p>
      ) : sortedMovies.length === 0 ? (
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
                <th
                  scope="col"
                  className="w-1/6 cursor-pointer px-3 py-3"
                  onClick={() => handleSort("user_rating")}
                >
                  <div className="flex items-center gap-2">
                    User Rating {renderSortIcon("user_rating")}
                  </div>
                </th>
                <th scope="col" className="w-1/6 px-3 py-3">
                  Genres
                </th>
                <th
                  scope="col"
                  className="w-1/5 cursor-pointer px-3 py-3"
                  onClick={() => handleSort("release_date")}
                >
                  <div className="flex items-center gap-2">
                    Release Date {renderSortIcon("release_date")}
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
                        src={`https://image.tmdb.org/t/p/w400/${movie.image}`}
                        alt={movie.title}
                        className="mr-4 h-30 min-w-20 object-cover transition-transform hover:scale-105"
                      />
                    </Link>
                  </td>
                  <td className="px-3 py-3 whitespace-nowrap">
                    <Link to={`/movies/${movie.id}`} className="text-blue-600 hover:underline">
                      {movie.title}
                    </Link>
                  </td>
                  <td className="px-3 py-3 whitespace-nowrap">
                    <UserRatingStar rating={movie.user_rating} />
                  </td>
                  <td className="px-3 py-3 whitespace-nowrap">{renderGenres(movie.genres)}</td>
                  <td className="px-3 py-3 whitespace-nowrap">{movie.release_date}</td>
                  <td className="px-3 py-3 whitespace-nowrap">{movie.mpaa_rating}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
