package router

import (
	"log/slog"

	"bookService/src/database/initdb"
	"bookService/src/server/handlers/createbook"
	"bookService/src/server/middleware/loggining"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(log *slog.Logger, storage *initdb.DB) *chi.Mux {
	Router := chi.NewRouter()

	Router.Use(middleware.RealIP)
	Router.Use(middleware.RequestID)
	Router.Use(loggining.Unable(log))

	Router.Get("/add_book", createbook.New(log, storage))

	return Router
}
