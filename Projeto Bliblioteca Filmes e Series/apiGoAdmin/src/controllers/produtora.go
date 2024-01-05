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



func GetProdutoras(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase();
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
	db := database.GetDatabase();
	vars := mux.Vars(r)
	username := vars["username"]

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

func CreateNewProdutora(w http.ResponseWriter, r *http.Request){
	// Supondo que você tenha uma instância de DB chamada "db"
    var novaProdutora models.Produtora
	db := database.GetDatabase();
	if err := json.NewDecoder(r.Body).Decode(&novaProdutora); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Lógica para salvar o novo usuário no banco de dados usando GORM
    if err := db.Create(&novaProdutora).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Resposta com o novo usuário criado
    novaProdutoraResponse, _ := json.Marshal(novaProdutora)
    w.WriteHeader(http.StatusCreated)
    w.Write(novaProdutoraResponse)
}

func DeleteProdutora(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    username := vars["username"]
	db := database.GetDatabase();
    // Lógica para deletar o usuário no banco de dados usando GORM
    var produtora models.Produtora
    if err := db.Where("username = ?", username).Delete(&produtora).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Simulação de uma resposta
    resposta := fmt.Sprintf("Produtora com Username %s deletado com sucesso", username)
    w.Write([]byte(resposta))
}

func UpdateProdutora(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]
    db := database.GetDatabase()

    var produtoraAtualizado models.Produtora
    if err := json.NewDecoder(r.Body).Decode(&produtoraAtualizado); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Pesquisa pela produtora usando o username
    var produtoraExistente models.Produtora
    if err := db.Where("username = ?", username).First(&produtoraExistente).Error; err != nil {
        http.Error(w, fmt.Sprintf("Produtora com username %s não encontrada", username), http.StatusNotFound)
        return
    }

    // Verifica se a produtoraExistente não é nula
    if produtoraExistente.ID == 0 {
        http.Error(w, fmt.Sprintf("Produtora com username %s não encontrada", username), http.StatusNotFound)
        return
    }

    // Atualiza apenas os campos fornecidos, evitando campos em branco
    if err := db.Model(&produtoraExistente).Updates(&produtoraAtualizado).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Resposta com a produtora atualizada
    produtoraAtualizadoResponse, _ := json.Marshal(produtoraExistente)
    w.Write(produtoraAtualizadoResponse)
}
