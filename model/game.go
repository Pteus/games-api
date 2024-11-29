package model

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type Game struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Platforms []string  `json:"platforms"`
	Genre     string    `json:"genre"`
}

func (o *Game) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create")
}

func (o *Game) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List")
}

func (o *Game) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetByID")
}

func (o *Game) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateByID")
}

func (o *Game) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteByID")
}
