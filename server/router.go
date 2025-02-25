package server

import (
	"github.com/gorilla/mux"
)

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/upload", UploadFileHandler).Methods("POST")
	r.HandleFunc("/query", QueryHandler).Methods("POST")
}
