package controllers

import (
	"apiGoAdmin/src/database"
	"apiGoAdmin/src/models"
	"apiGoAdmin/src/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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


func SaveFilme(w http.ResponseWriter, r *http.Request){
    var novafilme models.Filme
	db := database.GetDatabase();
	if err := json.NewDecoder(r.Body).Decode(&novafilme); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Lógica para salvar o novo filme no banco de dados usando GORM
    if err := db.Create(&novafilme).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Resposta com o novo filme criado
    novafilmeResponse, _ := json.Marshal(novafilme)
    w.WriteHeader(http.StatusCreated)
    w.Write(novafilmeResponse)
}

func DeleteFilme(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    id := vars["id"]
	db := database.GetDatabase();
    // Lógica para deletar filme no banco de dados usando GORM
    var filme models.Filme
    if err := db.Where("id = ?", id).Delete(&filme).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Simulação de uma resposta
    resposta := fmt.Sprintf("filme com Id %s deletado com sucesso", id)
    w.Write([]byte(resposta))
}

func UpdateFilme(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    db := database.GetDatabase()

    var filmeAtualizado models.Filme
    if err := json.NewDecoder(r.Body).Decode(&filmeAtualizado); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Pesquisa pela filme usando o id
    var filmeExistente models.Filme
    if err := db.Where("id = ?", id).First(&filmeExistente).Error; err != nil {
        http.Error(w, fmt.Sprintf("filme com id %s não encontrado", id), http.StatusNotFound)
        return
    }

    // Verifica se a filmeExistente não é nula
    if filmeExistente.ID == 0 {
        http.Error(w, fmt.Sprintf("filme com id %s não encontrada", id), http.StatusNotFound)
        return
    }

    // Atualiza apenas os campos fornecidos, evitando campos em branco
    if err := db.Model(&filmeExistente).Updates(&filmeAtualizado).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Resposta com a filme atualizada
    filmeAtualizadoResponse, _ := json.Marshal(filmeExistente)
    w.Write(filmeAtualizadoResponse)
}