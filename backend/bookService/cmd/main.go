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

	err = db.Migrate(&models.Book{}, &models.User{}, &models.File{}, &models.Comment{})
	if err != nil {
		logger.Error("Failed to migrate database", sl.Err(err))
		os.Exit(1)
	}

	BookRepository := repository.NewBookRepository(db.DB)
	CommentRepository := repository.NewCommentRepository(db.DB)

	bookService := service.NewBookService(BookRepository, logger)  // Создание сервиса для книг
	commentService := service.NewCommentService(CommentRepository) // Создание сервиса для комментариев

	// Инициализация хендлеров
	bookHandler := handler.NewBookHandler(bookService)          // Создание хендлера для книг
	commentHandler := handler.NewCommentHandler(commentService) // Создание хендлера для комментариев

	// Настройка маршрутизации
	router := server.SetupRouter(bookHandler, commentHandler)

	// Запуск сервера
	server.ListenServer(router, logger, cfg) // Запуск сервера
}
