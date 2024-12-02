package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/langdag/poe_chat_api/handlers"
)

func SetupRoutes() http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	// Middleware (optional)
	router.Use(middleware.Logger)

	// Define routes
	router.Post("/login", handlers.LoginHandler)
	router.Post("/registration", handlers.RegistrationHandler)
	router.Get("/", handlers.HomeHandler)

	return router
}
