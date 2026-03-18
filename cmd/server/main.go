package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/a-h/templ"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spoik/go-htmx-todo/internal/templates"
	"github.com/spoik/go-htmx-todo/internal/db"
)

func main() {
	db := db.Connect()
	defer db.Close()

	mux := createServer()
	startServer(mux, 8080)
}

func createServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(templates.Hello("World")))
	return mux
}

func startServer(mux *http.ServeMux, port int) {
	log.Printf("Starting server on :%d\n", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		mux,
	)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
