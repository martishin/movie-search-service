-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, picture_url, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM
    users
WHERE
    id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM
    users
WHERE
    email = $1;

-- name: ListUsers :many
SELECT *
FROM
    users
ORDER BY
    id;

-- name: DeleteUser :exec
DELETE
FROM
    users
WHERE
    id = $1;

-- name: LikeMovie :exec
INSERT INTO users_like_movies (user_id, movie_id)
VALUES ($1, $2)
ON CONFLICT (user_id, movie_id) DO NOTHING;

-- name: UnlikeMovie :exec
DELETE
FROM
    users_like_movies
WHERE
      user_id = $1
  AND movie_id = $2;

-- name: GetLikedMoviesByUser :many
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
    m.title;
