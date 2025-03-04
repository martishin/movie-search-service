ALTER TABLE movies
ADD COLUMN video VARCHAR(255);

-- Update existing movies with sample video URLs
UPDATE movies
SET video = CASE
    WHEN title = 'Highlander' THEN 'https://www.youtube.com/embed/omOZyLmNMJs?si=ODpSuvNLMS7HKoAR'
    WHEN title = 'Raiders of the Lost Ark' THEN 'https://www.youtube.com/embed/0xQSIdSRlAk?si=qkjzh3JzTJbc3Vab'
    WHEN title = 'The Godfather' THEN 'https://www.youtube.com/embed/UaVTIH8mujA?si=ZAhplY_EwUqvF-9x'
    ELSE NULL
END;
