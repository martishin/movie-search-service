CREATE TABLE movies_genres (
    id       SERIAL PRIMARY KEY,
    movie_id INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    CONSTRAINT fk_movies FOREIGN KEY (movie_id) REFERENCES movies (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_genres FOREIGN KEY (genre_id) REFERENCES genres (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Insert sample data
INSERT INTO movies_genres (movie_id, genre_id)
VALUES
    -- Highlander
    (1, 5),  -- Action
    (1, 12), -- Fantasy

    -- Raiders of the Lost Ark
   (2, 5),  -- Action
   (2, 11), -- Adventure

    -- The Godfather
   (3, 9),  -- Crime
   (3, 7),  -- Drama

    -- The Shawshank Redemption
    (4, 7),  -- Drama
    (4, 9),  -- Crime

    -- The Dark Knight
    (5, 5),  -- Action
    (5, 7),  -- Drama
    (5, 6),  -- Thriller
    (5, 13), -- Superhero

    -- Inception
    (6, 5),  -- Action
    (6, 2),  -- Sci-Fi
    (6, 6),  -- Thriller
    (6, 8),  -- Mystery

    -- Fight Club
    (7, 7),  -- Drama
    (7, 6),  -- Thriller

    -- Pulp Fiction
    (8, 9),  -- Crime
    (8, 7),  -- Drama

    -- Forrest Gump
    (9, 7),  -- Drama
    (9, 4),  -- Romance
    (9, 1),  -- Comedy

    -- The Matrix
    (10, 5),  -- Action
    (10, 2),  -- Sci-Fi

    -- The Lord of the Rings: The Fellowship of the Ring
    (11, 11), -- Adventure
    (11, 12), -- Fantasy
    (11, 7),  -- Drama

    -- The Lord of the Rings: The Two Towers
    (12, 11), -- Adventure
    (12, 12), -- Fantasy
    (12, 7),  -- Drama

    -- The Lord of the Rings: The Return of the King
    (13, 11), -- Adventure
    (13, 12), -- Fantasy
    (13, 7),  -- Drama

    -- Interstellar
    (14, 2),  -- Sci-Fi
    (14, 7),  -- Drama
    (14, 11), -- Adventure

    -- Gladiator
    (15, 5),  -- Action
    (15, 7),  -- Drama
    (15, 11), -- Adventure

    -- The Lion King
    (16, 10), -- Animation
    (16, 12), -- Fantasy
    (16, 7),  -- Drama

    -- Saving Private Ryan
    (17, 5),  -- Action
    (17, 7),  -- Drama
    (17, 11), -- Adventure

    -- Schindler's List
    (18, 7),  -- Drama
    (18, 9),  -- Crime

    -- The Silence of the Lambs
    (19, 6),  -- Thriller
    (19, 9),  -- Crime
    (19, 8),  -- Mystery

    -- Se7en
    (20, 6),  -- Thriller
    (20, 9),  -- Crime
    (20, 8);  -- Mystery
