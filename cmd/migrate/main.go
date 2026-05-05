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
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
