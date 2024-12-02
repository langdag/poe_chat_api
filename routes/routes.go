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
	router.Get("/", handlers.HomeHandler)                    // Example: GET /
	router.Get("/users", handlers.GetUsersHandler)           // Example: GET /users
	router.Post("/users", handlers.CreateUserHandler)        // Example: POST /users
	router.Get("/users/{id}", handlers.GetUserByIDHandler)   // Example: GET /users/:id
	router.Put("/users/{id}", handlers.UpdateUserHandler)    // Example: PUT /users/:id
	router.Delete("/users/{id}", handlers.DeleteUserHandler) // Example: DELETE /users/:id

	return router
}
