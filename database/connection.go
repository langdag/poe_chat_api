// database/connection.go
package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbpool *pgxpool.Pool // Global variable for the connection pool

// Connection initializes the database connection pool
func Connection() {
	// Database connection string (DSN)
	dsn := os.Getenv("DSN")

	// Initialize a connection pool
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse DSN: %v", err)
	}

	dbpool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	log.Println("Successfully connected to the database")
}

// GetDBPool returns the global database connection pool
func GetDBPool() *pgxpool.Pool {
	if dbpool == nil {
		log.Fatal("Database connection pool is not initialized.")
	}
	return dbpool
}
