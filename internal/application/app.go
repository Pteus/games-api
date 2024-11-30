package application

import (
	"fmt"
	"github.com/pteus/games-api/internal/handler"
	"github.com/pteus/games-api/internal/middleware"
	"github.com/pteus/games-api/internal/repository"
	"github.com/pteus/games-api/internal/routes"
	"log"
	"net/http"
)

type App struct {
	router http.Handler
	db     repository.GameRepository
}

func NewApp() *App {
	gameRepo, err := repository.NewPostgresGameRepository()
	if err != nil {
		log.Fatal(err)
	}
	gameHandler := &handler.GameHandler{Repo: gameRepo}

	return &App{
		router: routes.LoadRoutes(gameHandler),
	}
}

func (a *App) Start() error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.ApplyMiddleware(a.router, middleware.Logging, middleware.SetJSONContentType),
	}

	fmt.Println("Starting server on :8080")

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
