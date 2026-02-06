package main

import (
	"log"
	"net/http"

	"github.com/archis99/file-ms/internal/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.Load()

	// Router
	r := chi.NewRouter()

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Printf("File microservice started on port %s", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
