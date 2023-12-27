package controllers

import (
	"apiGoFilms/src/database"
	"apiGoFilms/src/models"
	"apiGoFilms/src/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProdutoras(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	var produtoras []models.Produtora
	db.Find(&produtoras)
	produtorasResponse, err := json.Marshal(produtoras)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(produtorasResponse)
	w.WriteHeader(http.StatusOK)
}

func GetProdutora(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	db := database.GetDatabase()
	var produtora models.Produtora
	result := db.Where("username = ?", username).First(&produtora)
	if result.Error != nil {
		log.Fatal(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		message := services.ResponseMessage("Não foi possível encontrar uma produtora com o username: " + username)
		notfound, _ := json.Marshal(message)
		w.Write(notfound)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	produtoraResponse, err := json.Marshal(produtora)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(produtoraResponse)
	w.WriteHeader(http.StatusOK)
}
