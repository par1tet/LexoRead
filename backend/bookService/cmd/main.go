package main

import (
	"bookService/src/config"
	"bookService/src/database/initdb"
	"bookService/src/database/models"
	"bookService/src/handler"
	"bookService/src/lib/prettylog"
	"bookService/src/lib/sl"
	"bookService/src/repository"
	"bookService/src/server"
	"bookService/src/service"
	"log/slog"
	"os"
)

func main() {

	cfg := config.MustLoad()

	logger := prettylog.NewLogger(slog.LevelDebug, true)

	db, err := initdb.Init(cfg.GetStorageDSN())
	if err != nil {
		logger.Error("Failed to connect to database", sl.Err(err))
		os.Exit(1)
	}

	err = db.Migrate(&models.Book{}, &models.File{}, &models.Comment{})
	if err != nil {
		logger.Error("Failed to migrate database", sl.Err(err))
		os.Exit(1)
	}

	BookRepository := repository.NewBookRepository(db.DB)
	CommentRepository := repository.NewCommentRepository(db.DB)
	RedisRepository := repository.NewRedisRepository("localhost:6379")

	bookService := service.NewBookService(BookRepository, logger)
	commentService := service.NewCommentService(CommentRepository)
	RedisService := service.NewRedisService(RedisRepository)

	bookHandler := handler.NewBookHandler(bookService)
	commentHandler := handler.NewCommentHandler(commentService)
	redisHandler := handler.NewRedisHandler(RedisService)

	router := server.SetupRouter(bookHandler, commentHandler, redisHandler)

	server.ListenServer(router, logger, cfg)
}
