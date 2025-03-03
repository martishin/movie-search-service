-- Add user_rating column to movies table
ALTER TABLE movies
    ADD COLUMN user_rating DECIMAL(2, 1) CHECK (user_rating BETWEEN 0 AND 5) DEFAULT NULL;

-- Update sample data with default ratings
UPDATE movies
SET
    user_rating = 4.5
WHERE
    title = 'Highlander';

UPDATE movies
SET
    user_rating = 4.8
WHERE
    title = 'Raiders of the Lost Ark';

UPDATE movies
SET
    user_rating = 4.9
WHERE
    title = 'The Godfather';
