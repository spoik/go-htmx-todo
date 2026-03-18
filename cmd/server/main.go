package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spoik/go-htmx-todo/internal/templates"
)

func main() {
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
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error: Could not ping database: %v\n", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(templates.Hello("World")))

	port := 8080

	log.Printf("Starting server on :%d\n", port)
	err = http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		mux,
	)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
