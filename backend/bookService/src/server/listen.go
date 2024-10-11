// src/server/listen.go
package server

import (
	"bookService/src/config"
	"bookService/src/lib/sl"
	"log/slog"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
)

func ListenServer(router *chi.Mux, slogger *slog.Logger, cfg *config.Config) {
	srv := &http.Server{
		Addr:         "localhost:" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	slogger.Info("Starting server on " + srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		slogger.Error("error on start server", sl.Err(err))
		os.Exit(1)
	}
}
