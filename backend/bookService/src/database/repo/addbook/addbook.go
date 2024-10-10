package addbook

import (
	"bookService/src/database/initdb"
	"bookService/src/database/models"
)

func Add(book *models.Book, db *initdb.DB) (*models.Book, error) {
	err := db.DB.Create(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}
