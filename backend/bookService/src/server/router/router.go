package router

import (
	"log/slog"

	"bookService/src/database/initdb"
	"bookService/src/server/handlers/createbook"
	"bookService/src/server/handlers/createcomment"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(log *slog.Logger, storage *initdb.DB) *chi.Mux {
	Router := chi.NewRouter()

	Router.Use(middleware.RealIP)
	Router.Use(middleware.RequestID)
	Router.Use(middleware.Logger)

	Router.Get("/add_book", createbook.New(log, storage))
	Router.Get("/add_comment", createcomment.New(log, storage))

	return Router
}
