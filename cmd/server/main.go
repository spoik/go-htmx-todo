package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spoik/go-htmx-todo/internal/db"
	"github.com/spoik/go-htmx-todo/internal/server"
)

func main() {
	db := db.Connect()
	defer db.Close()

	server.Start(8080)
}
