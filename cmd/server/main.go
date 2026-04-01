package main

import (
	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spoik/go-htmx-todo/internal/database"
	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server"
)

func main() {
	ctx := context.Background()

	dbCon := database.Connect(ctx)
	defer dbCon.Close(ctx)

	q := queries.New(dbCon)

	s := server.New(q)
	s.Start(8080)
}
