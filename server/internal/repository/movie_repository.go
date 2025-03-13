package repository

import (
	"context"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/martishin/movie-search-service/internal/db/generated"
	"github.com/martishin/movie-search-service/internal/model/domain"
)

type MovieRepository struct {
	pool    *pgxpool.Pool
	queries *db.Queries
}

func NewMovieRepository(postgresPool *pgxpool.Pool) *MovieRepository {
	return &MovieRepository{
		pool:    postgresPool,
		queries: db.New(postgresPool),
	}
}

func (r *MovieRepository) CreateMovie(ctx context.Context, movie domain.Movie) (db.Movie, error) {
	params := db.CreateMovieParams{
		Title:       movie.Title,
		ReleaseDate: pgtype.Date{Time: movie.ReleaseDate, Valid: true},
		Runtime:     pgtype.Int4{Int32: int32(movie.RunTime), Valid: true},
		MpaaRating:  pgtype.Text{String: movie.MPAARating, Valid: true},
		Description: pgtype.Text{String: movie.Description, Valid: true},
		Image:       pgtype.Text{String: movie.Image, Valid: true},
		Video:       pgtype.Text{String: movie.Video, Valid: true},
		UserRating:  pgtype.Numeric{Int: big.NewInt(int64(movie.UserRating * 10)), Exp: -1, Valid: true},
	}

	return r.queries.CreateMovie(ctx, params)
}

func (r *MovieRepository) GetMovieByID(ctx context.Context, id int) (db.Movie, error) {
	return r.queries.GetMovieByID(ctx, int32(id))
}

func (r *MovieRepository) ListMovies(ctx context.Context) ([]db.Movie, error) {
	return r.queries.ListMovies(ctx)
}

func (r *MovieRepository) UpdateMovie(ctx context.Context, movie domain.Movie) error {
	params := db.UpdateMovieParams{
		ID:          int32(movie.ID),
		Title:       movie.Title,
		ReleaseDate: pgtype.Date{Time: movie.ReleaseDate, Valid: true},
		Runtime:     pgtype.Int4{Int32: int32(movie.RunTime), Valid: true},
		MpaaRating:  pgtype.Text{String: movie.MPAARating, Valid: true},
		Description: pgtype.Text{String: movie.Description, Valid: true},
		Image:       pgtype.Text{String: movie.Image, Valid: true},
		Video:       pgtype.Text{String: movie.Image, Valid: true},
		UserRating:  pgtype.Numeric{Int: big.NewInt(int64(movie.UserRating * 10)), Exp: -1, Valid: true},
	}
	return r.queries.UpdateMovie(ctx, params)
}
func (r *MovieRepository) DeleteMovie(ctx context.Context, id int) error {
	return r.queries.DeleteMovie(ctx, int32(id))
}

func (r *MovieRepository) ListGenresByMovieID(ctx context.Context, movieID int) ([]db.Genre, error) {
	rows, err := r.queries.ListGenresByMovieID(ctx, int32(movieID))
	if err != nil {
		return nil, err
	}

	genres := make([]db.Genre, len(rows))
	for i, row := range rows {
		genres[i] = db.Genre{
			ID:    row.ID,
			Genre: row.Genre,
		}
	}
	return genres, nil
}

func (r *MovieRepository) AddMovieGenre(ctx context.Context, movieID, genreID int) error {
	params := db.AddMovieGenreParams{
		MovieID: int32(movieID),
		GenreID: int32(genreID),
	}

	return r.queries.AddMovieGenre(ctx, params)
}

func (r *MovieRepository) DeleteMovieGenres(ctx context.Context, movieID int) error {
	return r.queries.DeleteMovieGenres(ctx, int32(movieID))
}

func (r *MovieRepository) ListMoviesByGenre(ctx context.Context, genreID int) ([]db.Movie, error) {
	return r.queries.ListMoviesByGenre(ctx, int32(genreID))
}

func (r *MovieRepository) ListGenres(ctx context.Context) ([]db.Genre, error) {
	return r.queries.ListGenres(ctx)
}

func (r *MovieRepository) ListMoviesWithGenres(ctx context.Context) ([]db.ListMoviesWithGenresRow, error) {
	return r.queries.ListMoviesWithGenres(ctx)
}

func (r *MovieRepository) ListMoviesWithGenresAndLikes(ctx context.Context, userID int) ([]db.ListMoviesWithGenresAndLikeStatusRow, error) {
	return r.queries.ListMoviesWithGenresAndLikeStatus(ctx, int32(userID))
}

func (r *MovieRepository) IsMovieLikedByUser(ctx context.Context, movieID, userID int) (bool, error) {
	params := db.IsMovieLikedByUserParams{
		MovieID: int32(movieID),
		UserID:  int32(userID),
	}
	liked, err := r.queries.IsMovieLikedByUser(ctx, params)
	if err != nil {
		return false, err
	}
	return liked, nil
}

func (r *MovieRepository) GetLikedMovies(ctx context.Context, userID int) ([]db.GetLikedMoviesByUserRow, error) {
	return r.queries.GetLikedMoviesByUser(ctx, int32(userID))
}

func (r *MovieRepository) CreateMovieWithGenres(ctx context.Context, movie domain.Movie) (db.Movie, error) {
	// Start transaction
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return db.Movie{}, err
	}
	defer tx.Rollback(ctx)

	qtx := r.queries.WithTx(tx)

	// Step 1: Insert the movie
	params := db.CreateMovieParams{
		Title:       movie.Title,
		ReleaseDate: pgtype.Date{Time: movie.ReleaseDate, Valid: true},
		Runtime:     pgtype.Int4{Int32: int32(movie.RunTime), Valid: true},
		MpaaRating:  pgtype.Text{String: movie.MPAARating, Valid: true},
		Description: pgtype.Text{String: movie.Description, Valid: true},
		Image:       pgtype.Text{String: movie.Image, Valid: true},
		Video:       pgtype.Text{String: movie.Video, Valid: true},
		UserRating:  pgtype.Numeric{Int: big.NewInt(int64(movie.UserRating * 10)), Exp: -1, Valid: true},
	}

	dbMovie, err := qtx.CreateMovie(ctx, params)
	if err != nil {
		return db.Movie{}, err
	}

	// Step 2: Attach genres if provided
	genreIDs := getGenreIDs(movie.Genres)
	if len(genreIDs) > 0 {
		genreIDs32 := make([]int32, len(genreIDs))
		for i, id := range genreIDs {
			genreIDs32[i] = int32(id)
		}

		err = qtx.AttachGenresToMovie(ctx, db.AttachGenresToMovieParams{
			MovieID: dbMovie.ID,
			Column2: genreIDs32,
		})
		if err != nil {
			return db.Movie{}, err
		}
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return db.Movie{}, err
	}

	return dbMovie, nil
}

func getGenreIDs(genres []*domain.Genre) []int {
	var genreIDs []int
	for _, g := range genres {
		genreIDs = append(genreIDs, g.ID)
	}
	return genreIDs
}
