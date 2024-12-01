package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pteus/games-api/internal/model"
)

type PostgresGameRepository struct {
	db *sql.DB
}

func NewPostgresGameRepository() (*PostgresGameRepository, error) {
	connStr := "user=postgres dbname=games-api password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening DB connection: %v", err)
	}

	// Ping to ensure the connection is established
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection failed: %v", err)
	}

	return &PostgresGameRepository{db: db}, nil
}

func (r *PostgresGameRepository) GetAll() ([]model.Game, error) {
	var games []model.Game

	rows, err := r.db.Query("SELECT id, name, genre FROM games")
	if err != nil {
		return nil, fmt.Errorf("error fetching games: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var game model.Game
		if err := rows.Scan(&game.ID, &game.Name, &game.Genre); err != nil {
			return nil, fmt.Errorf("error scanning game: %v", err)
		}

		// Retrieve associated platforms
		platforms, err := r.getPlatformsByGameID(game.ID)
		if err != nil {
			return nil, fmt.Errorf("error fetching platforms for game %v: %v", game.ID, err)
		}
		game.Platforms = platforms

		games = append(games, game)
	}

	return games, nil
}

func (r *PostgresGameRepository) GetByID(id uuid.UUID) (*model.Game, error) {
	var game model.Game
	err := r.db.QueryRow("SELECT id, name, genre FROM games WHERE id = $1", id).Scan(&game.ID, &game.Name, &game.Genre)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("game not found")
		}
		return nil, fmt.Errorf("error fetching game: %v", err)
	}

	platforms, err := r.getPlatformsByGameID(game.ID)
	if err != nil {
		return nil, fmt.Errorf("error fetching platforms for game %v: %v", game.ID, err)
	}
	game.Platforms = platforms

	return &game, nil
}

func (repo *PostgresGameRepository) Create(game model.Game) (uuid.UUID, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return uuid.Nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Insert the game into the games table
	game.ID = uuid.New()
	_, err = tx.Exec(`INSERT INTO games (id, name, genre) VALUES ($1, $2, $3)`,
		game.ID, game.Name, game.Genre)
	if err != nil {
		return uuid.Nil, err
	}

	// Insert or retrieve platform IDs
	platformIDs := []int{}
	for _, platform := range game.Platforms {
		var platformID int
		err = tx.QueryRow(
			`INSERT INTO platforms (name) VALUES ($1)
			ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name
			RETURNING id`, platform).Scan(&platformID)
		if err != nil {
			return uuid.Nil, err
		}
		platformIDs = append(platformIDs, platformID)
	}

	// Link the game with platforms in game_platforms table
	for _, platformID := range platformIDs {
		_, err = tx.Exec(
			`INSERT INTO game_platforms (game_id, platform_id) VALUES ($1, $2)`,
			game.ID, platformID)
		if err != nil {
			return uuid.Nil, err
		}
	}

	return game.ID, nil
}

func (r *PostgresGameRepository) DeleteById(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM game_platforms WHERE game_id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting associated platforms for game %v: %v", id, err)
	}

	_, err = r.db.Exec("DELETE FROM games WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting game %v: %v", id, err)
	}

	return nil
}

func (r *PostgresGameRepository) UpdateById(id uuid.UUID, game model.Game) error {
	// Update the game in the 'games' table
	_, err := r.db.Exec(
		"UPDATE games SET name = $1, genre = $2 WHERE id = $3",
		game.Name, game.Genre, id,
	)
	if err != nil {
		return fmt.Errorf("error updating game: %v", err)
	}

	// Delete the old platform associations
	_, err = r.db.Exec("DELETE FROM game_platforms WHERE game_id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting old platforms for game %v: %v", id, err)
	}

	// Insert the updated platforms into the 'game_platforms' table
	for _, platform := range game.Platforms {
		var platformID int
		err := r.db.QueryRow("SELECT id FROM platforms WHERE name = $1", platform).Scan(&platformID)
		if errors.Is(err, sql.ErrNoRows) {
			// Platform doesn't exist, insert it
			err = r.db.QueryRow("INSERT INTO platforms (name) VALUES ($1) RETURNING id", platform).Scan(&platformID)
			if err != nil {
				return fmt.Errorf("error inserting platform: %v", err)
			}
		} else if err != nil {
			return fmt.Errorf("error checking platform: %v", err)
		}

		// Reassociate the game with the new platforms in the 'game_platforms' table
		_, err = r.db.Exec(
			"INSERT INTO game_platforms (game_id, platform_id) VALUES ($1, $2)",
			id, platformID,
		)
		if err != nil {
			return fmt.Errorf("error associating game with platform: %v", err)
		}
	}

	return nil
}

func (r *PostgresGameRepository) getPlatformsByGameID(gameID uuid.UUID) ([]string, error) {
	var platforms []string
	rows, err := r.db.Query(`
        SELECT p.name 
        FROM platforms p
        JOIN game_platforms gp ON p.id = gp.platform_id
        WHERE gp.game_id = $1`, gameID)
	if err != nil {
		return nil, fmt.Errorf("error fetching platforms for game %v: %v", gameID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var platform string
		if err := rows.Scan(&platform); err != nil {
			return nil, fmt.Errorf("error scanning platform: %v", err)
		}
		platforms = append(platforms, platform)
	}

	return platforms, nil
}
