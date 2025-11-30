package main

import (
	"log/slog"
	"os"
)

func main() {
	config := config{
		addr: ":8000",
		db : dbConfig{
			dsn : "host=localhost user=postgres password=Shipuden_2005 dbname=ecom-api sslmode=disable",
		},
	}
	api := application{
		config: config,
		// db : dbConfig{},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Error launching the server: ", "error", err)
		os.Exit(1)
	}
}
