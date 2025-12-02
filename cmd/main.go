package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/bruneldev/ecom-api/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	config := config{
		addr: ":8000",
		db : dbConfig{
			dsn : env.GetEnv("GOOSE_DBSTRING","host=localhost user=postgres password=Shipuden_2005 dbname=ecom-api sslmode=disable"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//database
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.db.dsn)
	if err != nil {
		slog.Error("Error while connecting to database", "dsn", config.db.dsn)
		panic(err)

		}
		defer conn.Close(ctx)
		logger.Info("Connected to database", "dsn", config.db.dsn)



	api := application{
		config: config,
		db: conn,
	}



	if err := api.run(api.mount()); err != nil {
		slog.Error("Error launching the server: ", "error", err)
		os.Exit(1)
	}
}
