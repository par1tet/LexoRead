package server

import (
	"bookService/src/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRouter(bookHandler *handler.BookHandler,
	commentHandler *handler.CommentHandler,
	redisHandler *handler.RedisHandler) *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	r.Route("/books", func(r chi.Router) {
		r.Post("/add", bookHandler.CreateBook)                // In docs
		r.Get("/limit={limit}", bookHandler.GetBooks)         // In docs
		r.Get("/id/{book_id}", bookHandler.GetBookByID)       // In docs
		r.Get("/search/{query}", bookHandler.SearchByKeyword) // In docs
		r.Get("/similar/{book_id}", bookHandler.SimilarBooks) // In docs
		r.Post("/like/{book_id}", bookHandler.LikeBook)       // In docs
		r.Post("/dislike/{book_id}", bookHandler.DislikeBook) // In docs
		r.Delete("/{book_id}", bookHandler.DeleteBook)        // In docs
		r.Put("/", bookHandler.UpdateBook)                    // In docs
	})

	r.Route("/comments", func(r chi.Router) {
		r.Post("/add", commentHandler.CreateComment)
		r.Get("/comments/book_id={book_id}", commentHandler.GetCommentsByBookID)
		r.Get("/comment/{id}", commentHandler.GetCommentByID)
	})

	r.Route("/redis_test", func(r chi.Router) {
		r.Post("/{type_book}/add", redisHandler.AddBook)
		r.Delete("/{type_book}/{id}", redisHandler.DeleteBook)
		r.Get("/{type_book}/{id}", redisHandler.GetBook)
		r.Get("/{type_book}/", redisHandler.GetBooks)
		r.Put("/{type_book}/{id}", redisHandler.UpdateBook)
	})

	return r
}
