package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/starmanv4/rag-backend/rag"
)

var (
	docStorage       = make(map[string]string)
	mu               sync.Mutex
	lastUploadedFile string
)

type QueryRequest struct {
	Query string `json:"query"`
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upload received...")

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	mu.Lock()
	docStorage[handler.Filename] = string(bytes)
	lastUploadedFile = handler.Filename
	mu.Unlock()

	fmt.Printf("File uploaded successfully: %s\n", handler.Filename)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "File uploaded successfully",
		"filename": handler.Filename,
	})
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Query received...")

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var req QueryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	filename := lastUploadedFile
	content, exists := docStorage[filename]
	mu.Unlock()

	if filename == "" || !exists {
		http.Error(w, "No file uploaded yet", http.StatusNotFound)
		return
	}

	fmt.Printf("Processing query for file: %s\n", filename)

	context := rag.RetrieveContext(content, req.Query)
	fmt.Println("Retrieved context:", context)

	response := rag.GenerateResponse(context, req.Query)
	fmt.Println("Generated response:", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"response": response,
	})
}
