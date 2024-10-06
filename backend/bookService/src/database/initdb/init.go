package initdb

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"bookService/src/database/models"
	"bookService/src/lib/sl"
)

type DB struct {
	DB *gorm.DB
}

func Init(dsn string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}

func (db *DB) Migrate(models ...interface{}) error {
	err := db.DB.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) CreateBook(book *models.Book, log *slog.Logger) {
	db.DB.AutoMigrate(&models.Book{})
	result := db.DB.Create(&book)
	if result.Error != nil {
		log.Error("Failed to create book", sl.Err(result.Error))
		return
	}
	log.Info("Book created",
		slog.Int("id", book.ID),
		slog.String("title", book.Title),
		slog.String("author", book.Author),
		slog.Int("likes", int(book.Likes)),
	)
}
