package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/starmanv4/rag-backend/rag"
)

var (
	docStorage = make(map[string]string) // In-memory storage for uploaded documents
	mu         sync.Mutex
)

type QueryRequest struct {
	Filename string `json:"filename"`
	Query    string `json:"query"`
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Max file size: 10MB

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Store file in memory
	mu.Lock()
	docStorage[handler.Filename] = string(bytes)
	mu.Unlock()

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	var req QueryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Retrieve document content
	mu.Lock()
	content, exists := docStorage[req.Filename]
	mu.Unlock()
	if !exists {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Retrieve relevant context
	context := rag.RetrieveContext(content, req.Query)

	// Generate AI response
	response := rag.GenerateResponse(context, req.Query)

	// Return response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"response": response})
}
