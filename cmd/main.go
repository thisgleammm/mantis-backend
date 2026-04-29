package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/thisgleammm/mantis-backend/internal/env"
)

// @title Mantis API
// @version 1.0
// @description This is the API server for Mantis.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host mantis-backend.fly.dev
// @BasePath /api/v1

func main() {
	ctx := context.Background()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cfg := config{
		addr: ":" + port,
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=root password=password dbname=mantis port=5432 sslmode=disable"),
		},
	}

	//logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("Connected to database", "dsn", cfg.db.dsn)

	api := &application{
		Config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server has failed to start", "error", err)
		os.Exit(1)
	}
}
