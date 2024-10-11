// src/repository/book_repo.go
package repository

import (
	"bookService/src/database/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetBooks() ([]models.Book, error)
	GetBookByID(book *models.Book, bookID int) error
	LikeBook(book *models.Book) error
	DislikeBook(book *models.Book) error
}

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepo{db: db}
}

func (r *bookRepo) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepo) GetBookByID(book *models.Book, bookID int) error {
	return r.db.Preload("Comments").First(&book, bookID).Error
}

func (r *bookRepo) GetBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepo) LikeBook(book *models.Book) error {
	return r.db.Model(book).Where("id = ?", book.ID).Update("likes", book.Likes+1).Error
}

func (r *bookRepo) DislikeBook(book *models.Book) error {
	return r.db.Model(book).Where("id = ?", book.ID).Update("dis_likes", book.DisLikes+1).Error
}
