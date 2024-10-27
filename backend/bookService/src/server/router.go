package server

import (
	"bookService/src/handler"
	redisHandler "bookService/src/handler/redis"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRouter(bookHandler *handler.BookHandler,
	commentHandler *handler.CommentHandler, redisHandler *redisHandler.RedisHandler,
	redisTest *handler.RedisHandler) *chi.Mux {

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
		r.Post("/add", bookHandler.CreateBook)
		r.Get("/limit={limit}", bookHandler.GetBooks)
		r.Get("/id/{book_id}", bookHandler.GetBookByID)
		r.Get("/search", bookHandler.SearchByKeyword)
		r.Post("/like/{book_id}", bookHandler.LikeBook)
		r.Post("/dislike/{book_id}", bookHandler.DislikeBook)
		r.Delete("/{book_id}", bookHandler.DeleteBook)
		r.Put("/", bookHandler.UpdateBook)
	})

	r.Route("/comments", func(r chi.Router) {
		r.Post("/add_comments", commentHandler.CreateComment)
		r.Get("/comments/{book_id}", commentHandler.GetCommentsByBookID)
	})
	r.Route("/redis_test/{type_book}", func(chi.Router) {
		r.Get("/{id}", redisTest.GetBook)
	})
	r.Route("/redis/top", func(r chi.Router) {
		r.Get("/", redisHandler.GetTopBooks)
		r.Post("/add", redisHandler.AddTopBook)
		r.Delete("/{id}", redisHandler.DeleteTopBook)
	})

	r.Route("/redis/default", func(r chi.Router) {

		r.Get("/{id}", redisHandler.GetBook)
		r.Post("/add", redisHandler.AddBook)
		r.Delete("/{id}", redisHandler.DeleteBook)
	})

	r.Route("/redis/business", func(r chi.Router) {
		r.Post("/add", redisHandler.AddBusinessBook)
		r.Get("/", redisHandler.GetBusinessBooks)
		r.Get("/{id}", redisHandler.GetBusinessBook)
		r.Delete("/{id}", redisHandler.DeleteBusinessBook)
	})

	r.Route("/redis/fiction", func(r chi.Router) {
		r.Post("/add", redisHandler.AddFictionBook)
		r.Get("/all", redisHandler.GetFictionBooks)
		r.Get("/book/{id}", redisHandler.GetFictionBook)
		r.Delete("/{id}", redisHandler.DeleteFictionBook)
	})

	r.Route("/redis/games", func(r chi.Router) {
		r.Post("/add", redisHandler.AddGamesBook)
		r.Get("/", redisHandler.GetGamesBooks)
		r.Get("/book/{id}", redisHandler.GetGamesBook)
		r.Delete("/{id}", redisHandler.DeleteGamesBook)
	})

	r.Route("/redis/literature", func(r chi.Router) {
		r.Post("/add", redisHandler.AddLiteratureBook)
		r.Get("/", redisHandler.GetLiteratureBooks)
		r.Get("/{id}", redisHandler.GetLiteratureBook)
		r.Delete("/{id}", redisHandler.DeleteLiteratureBook)
	})

	r.Route("/redis/novelty", func(r chi.Router) {
		r.Post("/add", redisHandler.AddNoveltyBook)
		r.Get("/", redisHandler.GetNoveltyBooks)
		r.Get("/{id}", redisHandler.GetNoveltyBook)
		r.Delete("/{id}", redisHandler.DeleteNoveltyBook)
	})

	r.Route("/redis/popular", func(r chi.Router) {
		r.Post("/add", redisHandler.AddPopularBook)
		r.Get("/", redisHandler.GetPopularBooks)
		r.Get("/{id}", redisHandler.GetPopularBook)
		r.Delete("/{id}", redisHandler.DeletePopularBook)
	})

	r.Route("/redis/psychology", func(r chi.Router) {
		r.Post("/add", redisHandler.AddPsychologyBook)
		r.Get("/", redisHandler.GetPsychologyBooks)
		r.Get("/{id}", redisHandler.GetPsychologyBook)
		r.Delete("/{id}", redisHandler.DeletePsychologyBook)
	})

	r.Route("/redis/romance", func(r chi.Router) {
		r.Post("/add", redisHandler.AddRomanceBook)
		r.Get("/", redisHandler.GetRomanceBooks)
		r.Get("/{id}", redisHandler.GetRomanceBook)
		r.Delete("/{id}", redisHandler.DeleteRomanceBook)
	})

	r.Route("/redis/self-development", func(r chi.Router) {
		r.Post("/add", redisHandler.AddSelfDevelopmentBook)
		r.Get("/", redisHandler.GetSelfDevelopmentBooks)
		r.Get("/{id}", redisHandler.GetSelfDevelopmentBook)
		r.Delete("/{id}", redisHandler.DeleteSelfDevelopmentBook)
	})

	return r
}
