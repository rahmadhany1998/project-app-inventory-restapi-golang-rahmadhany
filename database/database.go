package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=S4n14h1972@ dbname=golang_inventory sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	// Test Connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database failed: %w", err)
	}
	return db, nil
}
