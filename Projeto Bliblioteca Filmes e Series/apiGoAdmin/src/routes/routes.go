package routes

import (
	"apiGoAdmin/src/controllers"

	"github.com/gorilla/mux"
)

func StartRoutes(r *mux.Router) {
	r.HandleFunc("/filmes", controllers.GetFilmes).Methods("GET")//pega todos os filmes
	r.HandleFunc("/filmes/{id}", controllers.GetFilme).Methods("GET")//pega filmes por id

    r.HandleFunc("/filmes", controllers.SaveFilme).Methods("POST")//cria novo filme
	r.HandleFunc("/filmes/{id}", controllers.DeleteFilme).Methods("DELETE")//deleta filme
    r.HandleFunc("/filmes/{id}", controllers.UpdateFilme).Methods("PATCH")//modifica filme

	r.HandleFunc("/series", controllers.GetSeries).Methods("GET")//pega todas as series
	r.HandleFunc("/series/{id}", controllers.GetSerie).Methods("GET")//pega series por id

	//r.HandleFunc("/series", controllers.SaveSerie).Methods("POST")//cria nova serie
	//r.HandleFunc("/series/{id}", controllers.DeleteSerie).Methods("DELETE")//deleta serie
    //r.HandleFunc("/series/{id}", controllers.UpdateSerie).Methods("PATCH")//modifica serie

    r.HandleFunc("/produtoras", controllers.GetProdutoras).Methods("GET")//pega todas as produtoras
	r.HandleFunc("/produtoras/{username}", controllers.GetProdutora).Methods("GET")//pega produtora por username

    r.HandleFunc("/produtoras", controllers.CreateNewProdutora).Methods("POST")//cria produtora
    r.HandleFunc("/produtoras/{username}", controllers.DeleteProdutora).Methods("DELETE")//deleta produtora
    r.HandleFunc("/produtoras/{username}", controllers.UpdateProdutora).Methods("PATCH")//modifica produtora
}
