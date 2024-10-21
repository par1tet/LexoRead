package main

import (
	"bookService/src/config"
	"bookService/src/database/initdb"
	"bookService/src/database/models"
	docs2 "bookService/src/docs"
	"bookService/src/handler"
	redis_handler "bookService/src/handler/redis"
	"bookService/src/lib/prettylog"
	"bookService/src/lib/sl"
	"bookService/src/repository"
	redis_repo "bookService/src/repository/redis"
	"bookService/src/server"
	"bookService/src/service"
	redis_service "bookService/src/service/redis"
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
	RedisRepository := redis_repo.NewRedisRepository("localhost:6379")

	bookService := service.NewBookService(BookRepository, logger)
	commentService := service.NewCommentService(CommentRepository)
	RedisService := redis_service.NewRedisService(RedisRepository)

	bookHandler := handler.NewBookHandler(bookService)
	commentHandler := handler.NewCommentHandler(commentService)
	redisHandler := redis_handler.NewRedisHandler(RedisService)
	docs := docs2.GenDocs(bookHandler, redisHandler)
	router := server.SetupRouter(bookHandler, commentHandler, redisHandler)
	docs.SetupRoutes(router, docs)

	server.ListenServer(router, logger, cfg)
}
