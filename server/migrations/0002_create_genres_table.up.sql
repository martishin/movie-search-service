CREATE TABLE genres (
    id         SERIAL PRIMARY KEY,
    genre      VARCHAR(255) UNIQUE                 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TRIGGER set_timestamp_genres
    BEFORE UPDATE
    ON genres
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Insert sample data
INSERT INTO genres (genre, created_at, updated_at)
VALUES ('Comedy', NOW(), NOW()),
       ('Sci-Fi', NOW(), NOW()),
       ('Horror', NOW(), NOW()),
       ('Romance', NOW(), NOW()),
       ('Action', NOW(), NOW()),
       ('Thriller', NOW(), NOW()),
       ('Drama', NOW(), NOW()),
       ('Mystery', NOW(), NOW()),
       ('Crime', NOW(), NOW()),
       ('Animation', NOW(), NOW()),
       ('Adventure', NOW(), NOW()),
       ('Fantasy', NOW(), NOW()),
       ('Superhero', NOW(), NOW());
