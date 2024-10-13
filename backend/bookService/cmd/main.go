// @title LexoRead Book Service API
// @version 1.0
// @description API для работы с книгами в LexoRead.
// @termsOfService http://swagger.io/terms/

// @contact.name Support Team
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /
package main

import (
	"bookService/src/config"
	"bookService/src/database/initdb"
	"bookService/src/database/models"
	"bookService/src/handler"
	"bookService/src/handler/redis"
	"bookService/src/lib/prettylog"
	"bookService/src/lib/sl"
	"bookService/src/repository"
	"bookService/src/repository/redis"
	"bookService/src/server"
	"bookService/src/service"
	"bookService/src/service/redis"
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

	router := server.SetupRouter(bookHandler, commentHandler, redisHandler)

	server.ListenServer(router, logger, cfg)
}
