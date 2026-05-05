package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL is not defined in the environment variables")
	}

	// Check if the connection is real
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error when opening the database: %w", err)
	}

	return db, nil
}
