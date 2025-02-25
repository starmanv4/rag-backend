package main

import (
	"log"

	"github.com/starmanv4/rag-backend/rag"
	"github.com/starmanv4/rag-backend/server"
)

func main() {
	log.Println("Initializing OpenAI API...")
	rag.InitOpenAI() // Initialize OpenAI API

	log.Println("Starting server on port 8080...")
	server.StartServer()
}
