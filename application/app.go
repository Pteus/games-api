package application

import (
	"fmt"
	"games-api/middleware"
	"net/http"
)

type App struct {
	router http.Handler
}

func NewApp() *App {
	return &App{
		router: loadRoutes(),
	}
}

func (a *App) Start() error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(a.router),
	}

	fmt.Println("Starting server on :8080")

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
