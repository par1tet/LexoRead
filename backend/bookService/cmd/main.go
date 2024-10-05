package main

import (
	"bookService/src/config"
	"bookService/src/database/initdb"
	"bookService/src/lib/prettylog"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := prettylog.NewLogger(slog.LevelDebug, false)

	log.Info("Server started at cfg",
		slog.String("port", cfg.Port),
		slog.String("timeout", cfg.Timeout.String()),
		slog.String("idle_timeout", cfg.IdleTimeout.String()),
	)

	db, err := initdb.Init(cfg.GetStorageDSN())
	if err != nil {
		os.Exit(1)
	}
	_ = db

	// TODO: Init router

	// TODO: Run server

}
