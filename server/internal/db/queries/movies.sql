-- name: CreateMovie :one
INSERT INTO movies (title, release_date, runtime, mpaa_rating, description, image)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetMovieByID :one
SELECT *
FROM
    movies
WHERE
    id = $1;

-- name: ListMovies :many
SELECT *
FROM
    movies
ORDER BY
    id;

-- name: UpdateMovie :exec
UPDATE movies
SET title        = $2,
    release_date = $3,
    runtime      = $4,
    mpaa_rating  = $5,
    description  = $6,
    image        = $7
WHERE
    id = $1;

-- name: DeleteMovie :exec
DELETE
FROM
    movies
WHERE
    id = $1;

-- name: ListGenresByMovieID :many
SELECT
    g.id,
    g.genre
FROM
    movies_genres mg
    JOIN genres g ON mg.genre_id = g.id
WHERE
    mg.movie_id = $1;

-- name: AddMovieGenre :exec
INSERT INTO movies_genres (movie_id, genre_id)
VALUES ($1, $2);

-- name: DeleteMovieGenres :exec
DELETE
FROM
    movies_genres
WHERE
    movie_id = $1;

-- name: ListMoviesByGenre :many
SELECT
    m.*
FROM
    movies m
    JOIN movies_genres mg ON m.id = mg.movie_id
WHERE
    mg.genre_id = $1
ORDER BY
    m.title;

-- name: ListGenres :many
SELECT *
FROM
    genres
ORDER BY
    genre;
