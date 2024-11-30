package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pteus/games-api/internal/model"
	"github.com/pteus/games-api/internal/repository"
	"net/http"
)

type GameHandler struct {
	Repo repository.GameRepository
}

func (h *GameHandler) List(w http.ResponseWriter, r *http.Request) {
	games, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "Error getting all games", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		http.Error(w, "Error encoding games", http.StatusInternalServerError)
		return
	}
}

func (h *GameHandler) Create(w http.ResponseWriter, r *http.Request) {
	var game model.Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := h.Repo.Create(game)
	if err != nil {
		http.Error(w, "Error creating game", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Game created successfully",
		"id":      id.String(),
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (h *GameHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	game, err := h.Repo.GetByID(uuid.MustParse(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(game)
	if err != nil {
		http.Error(w, "Error encoding game", http.StatusInternalServerError)
	}
}

func (h *GameHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var updatedGame model.Game

	err := json.NewDecoder(r.Body).Decode(&updatedGame)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.Repo.UpdateById(uuid.MustParse(id), updatedGame)
	if err != nil {
		http.Error(w, "Error updating game", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *GameHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.Repo.DeleteById(uuid.MustParse(id))
	if err != nil {
		http.Error(w, "Error deleting game", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
