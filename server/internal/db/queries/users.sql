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
