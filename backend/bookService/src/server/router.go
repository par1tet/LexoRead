// src/server/router.go
package server

import (
	"bookService/src/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(bookHandler *handler.BookHandler,
	commentHandler *handler.CommentHandler, redisHandler *handler.RedisHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/add_book", bookHandler.CreateBook)
	r.Get("/books", bookHandler.GetBooks)
	r.Get("/books/{book_id}", bookHandler.GetBookByID)
	r.Post("/add_comments", commentHandler.CreateComment)
	r.Get("/comments", commentHandler.GetCommentsByBookID)
	r.Post("/like/{book_id}", bookHandler.LikeBook)
	r.Post("/dislike/{book_id}", bookHandler.DislikeBook)
	r.Post("/redis/add", redisHandler.AddBook)
	r.Get("/redis/{id}", redisHandler.GetBook)
	r.Get("/books/top", redisHandler.GetTopBooks)
	r.Delete("/redis/{id}", redisHandler.DeleteBook)

	return r
}
