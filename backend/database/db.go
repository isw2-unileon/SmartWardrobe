package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Fatal("DATABASE_URL no definida")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error conectando:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("No se puede hacer ping:", err)
	}

	DB = db
	log.Println("Conectado a Supabase")
}