package repository

import (
	"github.com/google/uuid"
	"github.com/pteus/games-api/internal/model"
)

type GameRepository interface {
	GetAll() ([]model.Game, error)
	GetByID(id uuid.UUID) (*model.Game, error)
	Create(game model.Game) error
	DeleteById(id uuid.UUID) error
	UpdateById(id uuid.UUID, game model.Game) error
}
