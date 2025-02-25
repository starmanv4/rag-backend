package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// StartServer initializes the server with proper CORS settings
func StartServer() {
	r := mux.NewRouter()

	// Define API routes
	setupRoutes(r)

	// CORS middleware using gorilla/handlers
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Allow frontend URL
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(), // Important for cookies/auth
	)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(r)))
}
