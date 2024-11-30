-- +goose Up
-- +goose StatementBegin
-- Enable the uuid-ossp extension for UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the games table
CREATE TABLE IF NOT EXISTS games (
                                     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                     name TEXT NOT NULL,
                                     genre TEXT NOT NULL
);

-- Create the platforms table
CREATE TABLE IF NOT EXISTS platforms (
                                         id SERIAL PRIMARY KEY,
                                         name TEXT NOT NULL UNIQUE
);

-- Create the many-to-many relationship between games and platforms
CREATE TABLE IF NOT EXISTS game_platforms (
                                              game_id UUID NOT NULL REFERENCES games(id) ON DELETE CASCADE,
                                              platform_id INT NOT NULL REFERENCES platforms(id) ON DELETE CASCADE,
                                              PRIMARY KEY (game_id, platform_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Disable the uuid-ossp extension (used to generate UUIDs)
DROP EXTENSION IF EXISTS "uuid-ossp";
-- Drop the game_platforms table
DROP TABLE IF EXISTS game_platforms;

-- Drop the platforms table
DROP TABLE IF EXISTS platforms;

-- Drop the games table
DROP TABLE IF EXISTS games;

-- +goose StatementEnd
