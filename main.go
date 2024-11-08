package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/langdag/poe_chat_api/config"
	"github.com/langdag/poe_chat_api/requests"
)

func main() {
	config.Connection()
	godotenv.Load(".env")

	serverPort := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL variable is not present")
	}

	if serverPort == "" {
		log.Fatal("PORT variable is not present")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/is_alive", requests.HandlerStatus)
	router.Mount("/v1", v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", serverPort),
	}

	log.Printf("Starting server on port %s", serverPort)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
