package model

import (
	"github.com/google/uuid"
)

type Game struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Platforms []string  `json:"platforms"`
	Genre     string    `json:"genre"`
}
