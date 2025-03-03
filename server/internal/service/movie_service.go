package service

import (
	"context"

	db "github.com/martishin/movie-search-service/internal/db/generated"
	"github.com/martishin/movie-search-service/internal/model/domain"
	"github.com/martishin/movie-search-service/internal/repository"
)

type MovieService struct {
	movieRepo *repository.MovieRepository
}

func NewMovieService(movieRepo *repository.MovieRepository) *MovieService {
	return &MovieService{movieRepo: movieRepo}
}

func (s *MovieService) CreateMovie(ctx context.Context, movie domain.Movie) (domain.Movie, error) {
	dbMovie, err := s.movieRepo.CreateMovie(ctx, movie)
	if err != nil {
		return domain.Movie{}, err
	}
	return mapDBMovieToDomainMovie(dbMovie), nil
}

func (s *MovieService) GetMovieByID(ctx context.Context, id int) (domain.Movie, error) {
	dbMovie, err := s.movieRepo.GetMovieByID(ctx, id)
	if err != nil {
		return domain.Movie{}, err
	}

	genres, err := s.movieRepo.ListGenresByMovieID(ctx, id)
	if err != nil {
		return domain.Movie{}, err
	}

	movie := mapDBMovieToDomainMovie(dbMovie)
	movie.Genres = mapDBGenresToDomainGenres(genres)
	return movie, nil
}

func (s *MovieService) ListMovies(ctx context.Context) ([]domain.Movie, error) {
	dbMovies, err := s.movieRepo.ListMovies(ctx)
	if err != nil {
		return nil, err
	}

	var movies []domain.Movie
	for _, dbMovie := range dbMovies {
		movies = append(movies, mapDBMovieToDomainMovie(dbMovie))
	}
	return movies, nil
}

func (s *MovieService) ListMoviesWithGenres(ctx context.Context) ([]domain.Movie, error) {
	rows, err := s.movieRepo.ListMoviesWithGenres(ctx)
	if err != nil {
		return nil, err
	}

	movieMap := make(map[int]domain.Movie)

	for _, row := range rows {
		movieID := int(row.MovieID)
		movie, exists := movieMap[movieID]
		userRating, _ := row.UserRating.Float64Value()
		if !exists {
			movie = domain.Movie{
				ID:          movieID,
				Title:       row.Title,
				ReleaseDate: row.ReleaseDate.Time,
				RunTime:     int(row.Runtime.Int32),
				MPAARating:  row.MpaaRating.String,
				Description: row.Description.String,
				Image:       row.Image.String,
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

	movies := make([]domain.Movie, 0, len(movieMap))
	for _, movie := range movieMap {
		movies = append(movies, movie)
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

func mapDBMovieToDomainMovie(dbMovie db.Movie) domain.Movie {
	userRating, _ := dbMovie.UserRating.Float64Value()

	return domain.Movie{
		ID:          int(dbMovie.ID),
		Title:       dbMovie.Title,
		ReleaseDate: dbMovie.ReleaseDate.Time,
		RunTime:     int(dbMovie.Runtime.Int32),
		MPAARating:  dbMovie.MpaaRating.String,
		Description: dbMovie.Description.String,
		Image:       dbMovie.Image.String,
		UserRating:  userRating.Float64,
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

func (s *MovieService) ListMoviesByGenre(ctx context.Context, genreID int) ([]domain.Movie, error) {
	dbMovies, err := s.movieRepo.ListMoviesByGenre(ctx, genreID)
	if err != nil {
		return nil, err
	}

	var movies []domain.Movie
	for _, dbMovie := range dbMovies {
		movies = append(movies, mapDBMovieToDomainMovie(dbMovie))
	}
	return movies, nil
}

func (s *MovieService) ListGenres(ctx context.Context) ([]domain.Genre, error) {
	dbGenres, err := s.movieRepo.ListGenres(ctx)
	if err != nil {
		return nil, err
	}

	var genres []domain.Genre
	for _, dbGenre := range dbGenres {
		genres = append(genres, domain.Genre{
			ID:    int(dbGenre.ID),
			Genre: dbGenre.Genre,
		})
	}
	return genres, nil
}
