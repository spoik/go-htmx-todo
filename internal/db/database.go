package db

import (
	"database/sql"
	"log"
	"os"
)

func Connect() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatalf(
			"DATABASE_URL environment variable is not set. This is required in order to connect to the database",
		)
	}

	db, err := sql.Open("pgx", dbURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Error: Could not ping database: %v\n", err)
	}

	return db
}
