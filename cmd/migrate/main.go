package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/thisgleammm/mantis-backend/internal/env"
)

func main() {
	dsn := env.RequiredString("GOOSE_DBSTRING")
	
	// Log connection attempt (safely)
	log.Printf("Connecting to database...")
	
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}
	log.Printf("Database connection successful")

	arguments := []string{}
	if len(os.Args) > 1 {
		arguments = os.Args[2:]
	}

	command := "up"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	if err := goose.Run(command, db, "internal/adapters/postgresql/migrations", arguments...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}
