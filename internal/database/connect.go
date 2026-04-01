package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) *pgx.Conn {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatalf(
			"DATABASE_URL environment variable is not set. This is required in order to connect to the database",
		)
	}

	con, err := pgx.Connect(ctx, dbURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	if err := con.Ping(ctx); err != nil {
		con.Close(ctx)
		log.Fatalf("Error: Could not ping database: %v\n", err)
	}

	return con
}
