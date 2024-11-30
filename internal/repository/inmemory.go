package repository

import (
	"errors"
	"github.com/pteus/games-api/internal/model"
	"sync"

	"github.com/google/uuid"
)

// InMemoryGameRepository implements GameRepository interface
type InMemoryGameRepository struct {
	games map[uuid.UUID]model.Game
	mu    sync.RWMutex // Protects access to the map
}

// NewInMemoryGameRepository creates a new instance of the in-memory repository
func NewInMemoryGameRepository() *InMemoryGameRepository {
	return &InMemoryGameRepository{
		games: make(map[uuid.UUID]model.Game),
	}
}

// GetAll retrieves all games from the repository
func (r *InMemoryGameRepository) GetAll() ([]model.Game, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	games := make([]model.Game, 0, len(r.games))
	for _, game := range r.games {
		games = append(games, game)
	}
	return games, nil
}

// GetByID retrieves a game by its ID
func (r *InMemoryGameRepository) GetByID(id uuid.UUID) (*model.Game, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	game, exists := r.games[id]
	if !exists {
		return nil, errors.New("game not found")
	}
	return &game, nil
}

// Create adds a new game to the repository, generating a new UUID for the game
func (r *InMemoryGameRepository) Create(game model.Game) (uuid.UUID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Generate a new UUID for the game
	game.ID = uuid.New()

	// Store the game in the repository
	r.games[game.ID] = game
	return game.ID, nil
}

// DeleteById removes a game from the repository by its ID
func (r *InMemoryGameRepository) DeleteById(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.games[id]; !exists {
		return errors.New("game not found")
	}
	delete(r.games, id)
	return nil
}

// UpdateById updates an existing game in the repository
func (r *InMemoryGameRepository) UpdateById(id uuid.UUID, updatedGame model.Game) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.games[id]; !exists {
		return errors.New("game not found")
	}

	// Update the game's fields while preserving the original ID
	updatedGame.ID = id
	r.games[id] = updatedGame
	return nil
}
