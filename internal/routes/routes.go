package routes

import (
	"github.com/pteus/games-api/internal/handler"
	"net/http"
)

func LoadRoutes(gameHandler *handler.GameHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /game", gameHandler.List)
	mux.HandleFunc("POST /game", gameHandler.Create)
	mux.HandleFunc("GET /game/{id}", gameHandler.GetByID)
	mux.HandleFunc("PUT /game/{id}", gameHandler.UpdateByID)
	mux.HandleFunc("DELETE /game/{id}", gameHandler.DeleteByID)

	return mux
}
