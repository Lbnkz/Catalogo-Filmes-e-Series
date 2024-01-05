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

func GetSeries(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	var series []models.Serie
	db.Find(&series)
	seriesResponse, err := json.Marshal(series)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(seriesResponse)
	w.WriteHeader(http.StatusOK)
}

func GetSerie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db := database.GetDatabase()
	var serie models.Serie
	result := db.Where("id = ?", id).First(&serie)
	if result.Error != nil {
		log.Fatal(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		message := services.ResponseMessage("Não foi possível encontrar uma série com o ID " + id)
		notfound, _ := json.Marshal(message)
		w.Write(notfound)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	serieResponse, err := json.Marshal(serie)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(serieResponse)
	w.WriteHeader(http.StatusOK)
}
