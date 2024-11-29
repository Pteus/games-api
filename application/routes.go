package application

import (
	"games-api/model"
	"net/http"
)

func loadRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	gameHandler := &model.Game{}

	mux.HandleFunc("GET /game", gameHandler.List)
	mux.HandleFunc("POST /game", gameHandler.Create)
	mux.HandleFunc("GET /game/{id}", gameHandler.GetByID)
	mux.HandleFunc("PUT /game/{id}", gameHandler.UpdateByID)
	mux.HandleFunc("DELETE /game/{id}", gameHandler.DeleteByID)

	return mux
}
