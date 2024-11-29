package main

import (
	"games-api/application"
	"log"
)

func main() {
	app := application.NewApp()

	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
