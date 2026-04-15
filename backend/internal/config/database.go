package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL is not defined in the environment variables")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error when opening the database: %w", err)
	}

	// Verificar si la conexión es real
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Error while ping the database: %w", err)
	}

	return db, nil
}
