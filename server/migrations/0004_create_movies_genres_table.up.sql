CREATE TABLE movies_genres (
    id       SERIAL PRIMARY KEY,
    movie_id INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    CONSTRAINT fk_movies FOREIGN KEY (movie_id) REFERENCES movies (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_genres FOREIGN KEY (genre_id) REFERENCES genres (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Insert sample data
INSERT INTO movies_genres (movie_id, genre_id)
VALUES (1, 5),  -- Highlander -> Action
       (1, 12), -- Highlander -> Fantasy
       (2, 5),  -- Raiders of the Lost Ark -> Action
       (2, 11), -- Raiders of the Lost Ark -> Adventure
       (3, 9),  -- The Godfather -> Crime
       (3, 7); -- The Godfather -> Drama
