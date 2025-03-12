package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	db "github.com/martishin/movie-search-service/internal/db/generated"
	"github.com/martishin/movie-search-service/internal/middleware"
	"github.com/martishin/movie-search-service/internal/model/domain"
	"github.com/martishin/movie-search-service/internal/repository"
	"github.com/redis/go-redis/v9"
)

type MovieService struct {
	movieRepo   *repository.MovieRepository
	redisClient *redis.Client
}

func NewMovieService(movieRepo *repository.MovieRepository, redisClient *redis.Client) *MovieService {
	return &MovieService{movieRepo: movieRepo, redisClient: redisClient}
}

func (s *MovieService) CreateMovie(ctx context.Context, movie domain.Movie) (*domain.Movie, error) {
	dbMovie, err := s.movieRepo.CreateMovie(ctx, movie)
	if err != nil {
		return nil, err
	}
	return mapDBMovieToDomainMovie(&dbMovie), nil
}

func (s *MovieService) GetMovieByIDWithGenres(ctx context.Context, id int) (*domain.Movie, error) {
	logger := middleware.GetLogger(ctx)

	// Check Redis cache
	cacheKey := fmt.Sprintf("movie:%d", id)

	cachedMovie, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err == nil && cachedMovie != "" {
		var movie domain.Movie
		if err := json.Unmarshal([]byte(cachedMovie), &movie); err == nil {
			logger.Info("Fetched movie from Redis", slog.Int("movie_id", id))
			return &movie, nil
		}
	}

	// Fetch from database
	dbMovie, err := s.movieRepo.GetMovieByID(ctx, id)
	if err != nil {
		return nil, err
	}

	genres, err := s.movieRepo.ListGenresByMovieID(ctx, id)
	if err != nil {
		return nil, err
	}

	movie := mapDBMovieToDomainMovie(&dbMovie)
	movie.Genres = mapDBGenresToDomainGenres(genres)

	// Store in Redis with a TTL of 10 minutes
	movieJSON, _ := json.Marshal(movie)
	err = s.redisClient.Set(ctx, cacheKey, string(movieJSON), 10*time.Minute).Err()
	if err != nil {
		logger.Error("Failed to store movie in Redis", slog.Any("error", err))
	}

	return movie, nil
}

func (s *MovieService) GetMovieByIDWithGenresAndLike(ctx context.Context, movieID int, userID int) (*domain.MovieWithLike, error) {
	dbMovie, err := s.movieRepo.GetMovieByID(ctx, movieID)
	if err != nil {
		return nil, err
	}

	genres, err := s.movieRepo.ListGenresByMovieID(ctx, movieID)
	if err != nil {
		return nil, err
	}

	isLiked, err := s.movieRepo.IsMovieLikedByUser(ctx, movieID, userID)
	movie := mapDBMovieToDomainMovieWithLike(&dbMovie, isLiked)
	movie.Genres = mapDBGenresToDomainGenres(genres)

	return movie, nil
}

func (s *MovieService) ListMoviesWithGenres(ctx context.Context) ([]*domain.Movie, error) {
	logger := middleware.GetLogger(ctx)

	// Check Redis cache
	cacheKey := "movies"

	cachedMovies, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err == nil && cachedMovies != "" {
		var movies []*domain.Movie
		if err := json.Unmarshal([]byte(cachedMovies), &movies); err == nil {
			logger.Info("Fetched movies from Redis")
			return movies, nil
		}
	}

	// Fetch from database
	rows, err := s.movieRepo.ListMoviesWithGenres(ctx)
	if err != nil {
		return nil, err
	}

	movieMap := make(map[int]*domain.Movie)

	for _, row := range rows {
		movieID := int(row.MovieID)
		movie, exists := movieMap[movieID]
		userRating, _ := row.UserRating.Float64Value()
		if !exists {
			movie = &domain.Movie{
				ID:          movieID,
				Title:       row.Title,
				ReleaseDate: row.ReleaseDate.Time,
				RunTime:     int(row.Runtime.Int32),
				MPAARating:  row.MpaaRating.String,
				Description: row.Description.String,
				Image:       row.Image.String,
				Video:       row.Video.String,
				Genres:      []*domain.Genre{},
				UserRating:  userRating.Float64,
			}
		}

		if row.GenreID.Valid {
			movie.Genres = append(movie.Genres, &domain.Genre{
				ID:    int(row.GenreID.Int32),
				Genre: row.Genre.String,
			})
		}

		movieMap[movieID] = movie
	}

	movies := make([]*domain.Movie, 0, len(movieMap))
	for _, movie := range movieMap {
		movies = append(movies, movie)
	}

	// Store in Redis with a TTL of 10 minutes
	movieJSON, _ := json.Marshal(movies)
	err = s.redisClient.Set(ctx, cacheKey, string(movieJSON), 10*time.Minute).Err()
	if err != nil {
		logger.Error("Failed to store movie in Redis", slog.Any("error", err))
	}

	return movies, nil
}

func (s *MovieService) UpdateMovie(ctx context.Context, movie domain.Movie) error {
	return s.movieRepo.UpdateMovie(ctx, movie)
}

func (s *MovieService) DeleteMovie(ctx context.Context, id int) error {
	return s.movieRepo.DeleteMovie(ctx, id)
}

func (s *MovieService) UpdateMovieGenres(ctx context.Context, movieID int, genreIDs []int) error {
	err := s.movieRepo.DeleteMovieGenres(ctx, movieID)
	if err != nil {
		return err
	}

	for _, genreID := range genreIDs {
		err = s.movieRepo.AddMovieGenre(ctx, movieID, genreID)
		if err != nil {
			return err
		}
	}
	return nil
}

func mapDBMovieToDomainMovie(dbMovie *db.Movie) *domain.Movie {
	userRating, _ := dbMovie.UserRating.Float64Value()

	return &domain.Movie{
		ID:          int(dbMovie.ID),
		Title:       dbMovie.Title,
		ReleaseDate: dbMovie.ReleaseDate.Time,
		RunTime:     int(dbMovie.Runtime.Int32),
		MPAARating:  dbMovie.MpaaRating.String,
		Description: dbMovie.Description.String,
		Image:       dbMovie.Image.String,
		Video:       dbMovie.Video.String,
		UserRating:  userRating.Float64,
	}
}

func mapDBMovieToDomainMovieWithLike(dbMovie *db.Movie, isLiked bool) *domain.MovieWithLike {
	userRating, _ := dbMovie.UserRating.Float64Value()

	return &domain.MovieWithLike{
		ID:          int(dbMovie.ID),
		Title:       dbMovie.Title,
		ReleaseDate: dbMovie.ReleaseDate.Time,
		RunTime:     int(dbMovie.Runtime.Int32),
		MPAARating:  dbMovie.MpaaRating.String,
		Description: dbMovie.Description.String,
		Image:       dbMovie.Image.String,
		Video:       dbMovie.Video.String,
		UserRating:  userRating.Float64,
		IsLiked:     isLiked,
	}
}

func mapDBGenresToDomainGenres(dbGenres []db.Genre) []*domain.Genre {
	var genres []*domain.Genre
	for _, dbGenre := range dbGenres {
		genres = append(genres, &domain.Genre{
			ID:    int(dbGenre.ID),
			Genre: dbGenre.Genre,
		})
	}
	return genres
}

func (s *MovieService) ListMoviesByGenre(ctx context.Context, genreID int) ([]*domain.Movie, error) {
	dbMovies, err := s.movieRepo.ListMoviesByGenre(ctx, genreID)
	if err != nil {
		return nil, err
	}

	var movies []*domain.Movie
	for _, dbMovie := range dbMovies {
		movies = append(movies, mapDBMovieToDomainMovie(&dbMovie))
	}
	return movies, nil
}

func (s *MovieService) ListGenres(ctx context.Context) ([]*domain.Genre, error) {
	dbGenres, err := s.movieRepo.ListGenres(ctx)
	if err != nil {
		return nil, err
	}

	var genres []*domain.Genre
	for _, dbGenre := range dbGenres {
		genres = append(genres, &domain.Genre{
			ID:    int(dbGenre.ID),
			Genre: dbGenre.Genre,
		})
	}
	return genres, nil
}

func (s *MovieService) ListMoviesWithGenresAndLikes(ctx context.Context, userID int) ([]*domain.MovieWithLike, error) {
	dbMovies, err := s.movieRepo.ListMoviesWithGenresAndLikes(ctx, userID)
	if err != nil {
		return nil, err
	}

	movieMap := make(map[int]*domain.MovieWithLike)

	for _, row := range dbMovies {
		movieID := int(row.MovieID)
		movie, exists := movieMap[movieID]
		userRating, _ := row.UserRating.Float64Value()

		if !exists {
			movie = &domain.MovieWithLike{
				ID:          movieID,
				Title:       row.Title,
				ReleaseDate: row.ReleaseDate.Time,
				RunTime:     int(row.Runtime.Int32),
				MPAARating:  row.MpaaRating.String,
				Description: row.Description.String,
				Image:       row.Image.String,
				Video:       row.Video.String,
				Genres:      []*domain.Genre{},
				UserRating:  userRating.Float64,
				IsLiked:     row.IsLiked,
			}
		}

		if row.GenreID.Valid {
			movie.Genres = append(movie.Genres, &domain.Genre{
				ID:    int(row.GenreID.Int32),
				Genre: row.Genre.String,
			})
		}

		movieMap[movieID] = movie
	}

	// Convert map to slice
	movies := make([]*domain.MovieWithLike, 0, len(movieMap))
	for _, movie := range movieMap {
		movies = append(movies, movie)
	}

	return movies, nil
}

func (s *MovieService) GetLikedMovies(ctx context.Context, userID int) ([]*domain.Movie, error) {
	dbMovies, err := s.movieRepo.GetLikedMovies(ctx, userID)
	if err != nil {
		return nil, err
	}

	movieMap := make(map[int]*domain.Movie)

	for _, row := range dbMovies {
		movieID := int(row.MovieID)
		movie, exists := movieMap[movieID]
		userRating, _ := row.UserRating.Float64Value()

		if !exists {
			movie = &domain.Movie{
				ID:          movieID,
				Title:       row.Title,
				ReleaseDate: row.ReleaseDate.Time,
				RunTime:     int(row.Runtime.Int32),
				MPAARating:  row.MpaaRating.String,
				Description: row.Description.String,
				Image:       row.Image.String,
				UserRating:  userRating.Float64,
				Video:       row.Video.String,
				Genres:      []*domain.Genre{},
			}
		}

		if row.GenreID.Valid {
			movie.Genres = append(movie.Genres, &domain.Genre{
				ID:    int(row.GenreID.Int32),
				Genre: row.Genre.String,
			})
		}

		movieMap[movieID] = movie
	}

	movies := make([]*domain.Movie, 0)
	for _, movie := range movieMap {
		movies = append(movies, movie)
	}

	return movies, nil
}
