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
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	r.Route("/books", func(r chi.Router) {
		r.Post("/add", bookHandler.CreateBook)
		r.Get("/", bookHandler.GetBooks)
		r.Get("/{book_id}", bookHandler.GetBookByID)
		r.Get("/search", bookHandler.SearchByKeyword)
		r.Post("/like/{book_id}", bookHandler.LikeBook)
		r.Post("/dislike/{book_id}", bookHandler.DislikeBook)
	})
	r.Route("/comments", func(r chi.Router) {
		r.Post("/add_comments", commentHandler.CreateComment)
		r.Get("/comments/{book_id}", commentHandler.GetCommentsByBookID)
	})
	r.Route("/redis", func(r chi.Router) {
		r.Post("/add", redisHandler.AddBook)
		r.Get("/{id}", redisHandler.GetBook)
		r.Get("/top_books", redisHandler.GetTopBooks)
		r.Delete("/redis/{id}", redisHandler.DeleteBook)
	})

	return r
}
