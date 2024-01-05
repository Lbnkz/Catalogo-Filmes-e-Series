package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"apiGoFilms/src/database"
	"apiGoFilms/src/models"
	"apiGoFilms/src/services"

	"github.com/gorilla/mux"
)

func GetFilmes(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	var filmes []models.Filme
	db.Find(&filmes)
	filmesResponse, err := json.Marshal(filmes)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(filmesResponse)
	w.WriteHeader(http.StatusOK)
}

func GetFilme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db := database.GetDatabase()
	var filme models.Filme
	result := db.Where("id = ?", id).First(&filme)
	if result.Error != nil {
		log.Fatal(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		message := services.ResponseMessage("Não foi possível encontrar um filme com o ID " + id)
		notfound, _ := json.Marshal(message)
		w.Write(notfound)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	filmeResponse, err := json.Marshal(filme)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(filmeResponse)
	w.WriteHeader(http.StatusOK)
}
