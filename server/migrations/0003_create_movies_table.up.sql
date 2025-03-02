CREATE TABLE movies (
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(512)                        NOT NULL,
    release_date DATE,
    runtime      INTEGER,
    mpaa_rating  VARCHAR(10),
    description  TEXT,
    image        VARCHAR(255),
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TRIGGER set_timestamp_movies
    BEFORE UPDATE
    ON movies
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Insert sample data
INSERT INTO movies (title, release_date, runtime, mpaa_rating, description, image, created_at, updated_at)
VALUES ('Highlander', '1986-03-07', 116, 'R',
        'He fought his first battle on the Scottish Highlands in 1536. He will fight his greatest battle on the streets of New York City in 1986. His name is Connor MacLeod. He is immortal.',
        '/8Z8dptJEypuLoOQro1WugD855YE.jpg', NOW(), NOW()),
       ('Raiders of the Lost Ark', '1981-06-12', 115, 'PG-13',
        'Archaeology professor Indiana Jones ventures to seize a biblical artefact known as the Ark of the Covenant. While doing so, he puts up a fight against Renee and a troop of Nazis.',
        '/ceG9VzoRAVGwivFU403Wc3AHRys.jpg', NOW(), NOW()),
       ('The Godfather', '1972-03-24', 175, '18A',
        'The aging patriarch of an organized crime dynasty in postwar New York City transfers control of his clandestine empire to his reluctant youngest son.',
        '/3bhkrj58Vtu7enYsRolD1fZdja1.jpg', NOW(), NOW());
