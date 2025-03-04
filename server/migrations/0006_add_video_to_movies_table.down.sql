ALTER TABLE movies
ADD COLUMN video VARCHAR(255);

-- Update existing movies with sample video URLs
UPDATE movies
SET video = CASE
    WHEN title = 'Highlander' THEN 'kedW1xO3Zbo'
    WHEN title = 'Raiders of the Lost Ark' THEN '0xQSIdSRlAk'
    WHEN title = 'The Godfather' THEN 'UaVTIH8mujA'
    ELSE NULL
END;
