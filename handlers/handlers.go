package handlers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"database/sql"
	"errors"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/langdag/poe_chat_api/database"
	"github.com/langdag/poe_chat_api/models"
	"github.com/langdag/poe_chat_api/requests"
	"github.com/langdag/poe_chat_api/validations"
)

var jwtKey = []byte(os.Getenv("JWT"))

type TokenResponse struct {
	Token string `json:"token"`
}

// GenerateJWT generates a JWT token for authenticated users
func GenerateJWT(user models.DefaultUser) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.MapClaims{
		"id":   user.ID,
		"exp": jwt.NewNumericDate(expirationTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error signing token: %v", err) // Log the error for debugging
		return "", err
	}
	return signedToken, nil
}

// LoginHandler handles user login and returns a JWT token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.LoginUser
	requests.ParseJSON(r, &user)

	validationsError := validations.HandleValidations(w, user)

	if validationsError != nil {
		return
	}

	var existingUser models.DefaultUser

	db := database.GetDBPool()

	query := `SELECT email, password FROM users WHERE email = $1 AND password = $2 LIMIT 1`
	err := db.QueryRow(context.Background(), query, user.Email, user.Password).Scan(&existingUser.Email, &existingUser.Password)

	if err != nil {
		requests.HandlerError(w, http.StatusNotFound, "Invalid email or password")
		return
	}

	token, err := GenerateJWT(existingUser)
	if err != nil {
		requests.HandlerError(w, http.StatusInternalServerError, "Error generating token")
		return
	}
	requests.HandlerResponse(w, http.StatusOK, TokenResponse{Token: token})
}

// RegistrationHandler handles user registration
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var user models.RegisterUser
	requests.ParseJSON(r, &user)

	validationsError := validations.HandleValidations(w, user)
	if validationsError != nil {
		return
	}

	db := database.GetDBPool()

	// Check if user with the given username or email already exists
	var existingUser models.DefaultUser

	checkQuery := `SELECT username FROM users WHERE username = $1 LIMIT 1`
	err := db.QueryRow(context.Background(), checkQuery, user.Username, user.Email).Scan(&existingUser.Username)
	if err != nil {
		requests.HandlerError(w, http.StatusConflict, "User with username or email already exists")
		return
	}

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`

	_, err = db.Exec(context.Background(), query, user.Username, user.Email, user.Password)

	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	response := requests.SuccessResponse{
		Message: "User registered successfully",
		Data:    user,
	}

	requests.HandlerResponse(w, http.StatusOK, response)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user models.DefaultUser

	db := database.GetDBPool()
	query := `SELECT * FROM users WHERE id = $1`
	err := db.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		requests.HandlerError(w, http.StatusNotFound, "User not found")
		return
	}
	if err != nil {
		requests.HandlerError(w, http.StatusInternalServerError, "Error fetching user")
		return
	}

	response := requests.SuccessResponse{
		Message: "User have being fetched successfully",
		Data:    user,
	}
	requests.HandlerResponse(w, http.StatusOK, response)
}

// Middleware to verify JWT tokens
func JWTAuthMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			requests.HandlerError(w, http.StatusUnauthorized, "Missing token")
			return
		}

		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			requests.HandlerError(w, http.StatusUnauthorized, "Access denied")
			return
		}

		// Token is valid; proceed to the next handler
		handlerFunc(w, r)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := requests.SuccessResponse{
		Message: "Welcome to the Home Page!",
		Data:    nil,
	}
	requests.HandlerResponse(w, http.StatusOK, response)
}

//Add comparing for token claims id with the user id from the database
//Refactor to use refresh token with access tokens
