CREATE TABLE users_like_movies (
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER                             NOT NULL,
    movie_id   INTEGER                             NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_movies FOREIGN KEY (movie_id) REFERENCES movies (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT unique_user_movie UNIQUE (user_id, movie_id) -- Ensures a user can like a movie only once
);
