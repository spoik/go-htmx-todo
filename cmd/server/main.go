package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spoik/go-htmx-todo/internal/database"
	"github.com/spoik/go-htmx-todo/internal/server"
)

func main() {
	db := database.Connect()
	defer db.Close()

	serv := server.Create()

	server.Start(serv, 8080)
}
