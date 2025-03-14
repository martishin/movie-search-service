package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/martishin/movie-search-service/internal/db/generated"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(postgresPool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		queries: db.New(postgresPool),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, firstName, lastName, email, pictureURL string, password string) (db.User, error) {
	params := db.CreateUserParams{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		PictureUrl: pgtype.Text{String: pictureURL, Valid: true},
		Password:   pgtype.Text{String: password, Valid: password != ""},
	}
	return r.queries.CreateUser(ctx, params)
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (db.User, error) {
	return r.queries.GetUserByID(ctx, int32(id))
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return r.queries.GetUserByEmail(ctx, email)
}

func (r *UserRepository) LikeMovie(ctx context.Context, userID, movieID int) error {
	return r.queries.LikeMovie(ctx, db.LikeMovieParams{
		UserID:  int32(userID),
		MovieID: int32(movieID),
	})
}

func (r *UserRepository) UnlikeMovie(ctx context.Context, userID, movieID int) error {
	return r.queries.UnlikeMovie(ctx, db.UnlikeMovieParams{
		UserID:  int32(userID),
		MovieID: int32(movieID),
	})
}
