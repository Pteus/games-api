package handler

import (
	"fmt"
	"github.com/pteus/games-api/internal/repository"
	"net/http"
)

type GameHandler struct {
	Db repository.GameRepository
}

func (o *GameHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List")
}

func (o *GameHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create")
}

func (o *GameHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetByID")
}

func (o *GameHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateByID")
}

func (o *GameHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteByID")
}
