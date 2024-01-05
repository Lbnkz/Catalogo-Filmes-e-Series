package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"apiGoAdmin/src/database"
	"apiGoAdmin/src/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.StartDB()
	var MyRouter = r.PathPrefix("/api").Subrouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	FinalRouter := c.Handler(r)
	routes.StartRoutes(MyRouter)
	PORT := os.Getenv("PORT")
	fmt.Println("Server running on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, FinalRouter))

}
