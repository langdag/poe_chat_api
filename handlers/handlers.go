package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/langdag/poe_chat_api/database"
)

var jwtKey = []byte(os.Getenv("JWT"))

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GenerateJWT generates a JWT token for authenticated users
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
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
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var existingUser User

	db := database.GetDBPool()

	query := `SELECT username, password FROM users WHERE username = $1 AND password = $2 LIMIT 1`
	err := db.QueryRow(context.Background(), query, user.Username, user.Password).Scan(&existingUser.Username, &existingUser.Password)

	if err != nil {
		w.Write([]byte("User not found!"))
		return
	}

	token, err := GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// RegistrationHandler handles user registration
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Username == "" {
		http.Error(w, "Username cannot be empty", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "Email cannot be empty", http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		http.Error(w, "Password cannot be empty", http.StatusBadRequest)
		return
	}

	db := database.GetDBPool()

	// Check if user with the given username or email already exists
	var existingUser User

	checkQuery := `SELECT username FROM users WHERE username = $1 LIMIT 1`
	checked_user := db.QueryRow(context.Background(), checkQuery, user.Username, user.Email).Scan(&existingUser.Username)
	if checked_user != nil {

		w.WriteHeader(http.StatusConflict) // HTTP 409 Conflict status code
		w.Write([]byte(fmt.Sprintf("User with username %s already exists", user.Username)))
		return
	}

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`

	_, err := db.Exec(context.Background(), query, user.Username, user.Email, user.Password)

	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

// AuthenticatedRoute demonstrates how to protect routes with JWT
func AuthenticatedRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected route"))
}

// Middleware to verify JWT tokens
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid; proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page!"))
}
