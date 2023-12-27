package routes

import (
	"apiGoFilms/src/controllers"

	"github.com/gorilla/mux"
)

func StartRoutes(r *mux.Router) {
	r.HandleFunc("/filmes", controllers.GetFilmes).Methods("GET")
	r.HandleFunc("/filmes/{id}", controllers.GetFilme).Methods("GET")

	r.HandleFunc("/series", controllers.GetSeries).Methods("GET")
	r.HandleFunc("/series/{id}", controllers.GetSerie).Methods("GET")

	r.HandleFunc("/produtoras", controllers.GetProdutoras).Methods("GET")
	r.HandleFunc("/produtoras/{username}", controllers.GetProdutora).Methods("GET")
}
