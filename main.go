package main

import (
	"github.com/pteus/games-api/internal/application"
	"log"
)

func main() {
	app := application.NewApp()

	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
