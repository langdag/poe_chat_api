package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/langdag/poe_chat_api/database"
	"github.com/langdag/poe_chat_api/routes"
)

func main() {
	godotenv.Load(".env")
	database.Connection()

	port := os.Getenv("PORT")
	if port == "" {
			port = "8080"
	}

	router := routes.SetupRoutes()

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
			log.Fatal(err)
	}
}
