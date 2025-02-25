package server

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-chi/cors"
)

func StartServer() {
	r := mux.NewRouter()

	// Enable CORS for frontend communication
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
	}))

	// Define API routes
	setupRoutes(r)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
