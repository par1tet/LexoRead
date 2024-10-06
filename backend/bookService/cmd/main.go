package main

import (
	"bookService/src/config"
	"bookService/src/database/initdb"
	"bookService/src/database/models"
	"bookService/src/lib/prettylog"
	"bookService/src/server/listen"
	"bookService/src/server/router"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := prettylog.NewLogger(slog.LevelInfo, false)

	log.Info("Server started at cfg",
		slog.String("port", cfg.Port),
		slog.String("timeout", cfg.Timeout.String()),
		slog.String("idle_timeout", cfg.IdleTimeout.String()),
	)

	db, err := initdb.Init(cfg.GetStorageDSN())
	if err != nil {
		os.Exit(1)
	}
	db.Migrate(&models.Book{}, &models.Comment{}, &models.File{})

	Router := router.New(log, db)

	listen.ListenServer(Router, log, cfg)

}
