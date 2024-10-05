package main

import (
	"bookService/src/config"
	"bookService/src/database/initdb"
	"bookService/src/database/models"
	"bookService/src/lib/prettylog"
	"bookService/src/lib/sl"
	"log/slog"
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
		log.Error("Failed to init db", sl.Err(err))
	}
	book := &models.Book{
		Title:  "Слово пацана",
		Author: "Пальто",
	}
	db.CreateBook(book, log)
	if err != nil {
		log.Error("Failed to create book", sl.Err(err))
		return
	}

	// Init router

	// Run server

}
