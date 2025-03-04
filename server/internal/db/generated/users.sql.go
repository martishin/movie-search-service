// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, picture_url, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, first_name, last_name, email, password, picture_url, created_at, updated_at
`

type CreateUserParams struct {
	FirstName  string
	LastName   string
	Email      string
	PictureUrl pgtype.Text
	Password   pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PictureUrl,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.PictureUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE
FROM
    users
WHERE
    id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getLikedMoviesByUser = `-- name: GetLikedMoviesByUser :many
SELECT
    m.id AS movie_id,
    m.title,
    m.release_date,
    m.runtime,
    m.mpaa_rating,
    m.description,
    m.image,
    m.user_rating,
    m.video,
    g.id AS genre_id,
    g.genre
FROM
    users_like_movies ulm
        JOIN movies m ON ulm.movie_id = m.id
        LEFT JOIN movies_genres mg ON m.id = mg.movie_id
        LEFT JOIN genres g ON mg.genre_id = g.id
WHERE
    ulm.user_id = $1
ORDER BY
    m.title
`

type GetLikedMoviesByUserRow struct {
	MovieID     int32
	Title       string
	ReleaseDate pgtype.Date
	Runtime     pgtype.Int4
	MpaaRating  pgtype.Text
	Description pgtype.Text
	Image       pgtype.Text
	UserRating  pgtype.Numeric
	Video       pgtype.Text
	GenreID     pgtype.Int4
	Genre       pgtype.Text
}

func (q *Queries) GetLikedMoviesByUser(ctx context.Context, userID int32) ([]GetLikedMoviesByUserRow, error) {
	rows, err := q.db.Query(ctx, getLikedMoviesByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLikedMoviesByUserRow
	for rows.Next() {
		var i GetLikedMoviesByUserRow
		if err := rows.Scan(
			&i.MovieID,
			&i.Title,
			&i.ReleaseDate,
			&i.Runtime,
			&i.MpaaRating,
			&i.Description,
			&i.Image,
			&i.UserRating,
			&i.Video,
			&i.GenreID,
			&i.Genre,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, password, picture_url, created_at, updated_at
FROM
    users
WHERE
    email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.PictureUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, first_name, last_name, email, password, picture_url, created_at, updated_at
FROM
    users
WHERE
    id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.PictureUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const likeMovie = `-- name: LikeMovie :exec
INSERT INTO users_like_movies (user_id, movie_id)
VALUES ($1, $2)
ON CONFLICT (user_id, movie_id) DO NOTHING
`

type LikeMovieParams struct {
	UserID  int32
	MovieID int32
}

func (q *Queries) LikeMovie(ctx context.Context, arg LikeMovieParams) error {
	_, err := q.db.Exec(ctx, likeMovie, arg.UserID, arg.MovieID)
	return err
}

const listUsers = `-- name: ListUsers :many
SELECT id, first_name, last_name, email, password, picture_url, created_at, updated_at
FROM
    users
ORDER BY
    id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.PictureUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unlikeMovie = `-- name: UnlikeMovie :exec
DELETE
FROM
    users_like_movies
WHERE
      user_id = $1
  AND movie_id = $2
`

type UnlikeMovieParams struct {
	UserID  int32
	MovieID int32
}

func (q *Queries) UnlikeMovie(ctx context.Context, arg UnlikeMovieParams) error {
	_, err := q.db.Exec(ctx, unlikeMovie, arg.UserID, arg.MovieID)
	return err
}
