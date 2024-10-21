package repository

import (
	"bookService/src/database/models"
	"errors"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetBooks(limit int) ([]models.Book, error)
	GetBookByID(book *models.Book, bookID int) error
	LikeBook(book_id string) error
	DislikeBook(book_id string) error
	SearchByKeyword(keyword string) ([]models.Book, error)
	DeleteBook(id string) error
	UpdateBook(book *models.Book) error
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
	return r.db.Preload("Comments").Preload("Files").First(&book, bookID).Error
}

func (r *bookRepo) GetBooks(limit int) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Limit(limit).
		Preload("Comments").
		Preload("Files").
		Find(&books).Error

	return books, err
}

func (r *bookRepo) LikeBook(book_id string) error {
	result := r.db.Model(models.Book{}).Where("id = ?",
		book_id).UpdateColumn("likes", gorm.Expr("likes + ?", 1))
	if result.RowsAffected == 0 {
		return errors.New("Book not found")
	}
	return result.Error
}

func (r *bookRepo) DislikeBook(book_id string) error {
	result := r.db.Model(models.Book{}).Where("id = ?",
		book_id).UpdateColumn("dis_likes", gorm.Expr("dis_likes + ?", 1))
	if result.RowsAffected == 0 {
		return errors.New("Book not found")
	}
	return result.Error
}

func (r *bookRepo) SearchByKeyword(keyword string) ([]models.Book, error) {
	var books []models.Book

	keyword = "%" + keyword + "%"
	result := r.db.Preload("Comments").Preload("Files").Where("title ILIKE ?",
		keyword).Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("по вашему запросу книги не найдены")
	}

	return books, nil
}

func (r *bookRepo) DeleteBook(id string) error {
	if err := r.db.First(&models.Book{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Book not found")
		}
		return err
	}
	return r.db.Delete(&models.Book{}, id).Error
}

func (r *bookRepo) UpdateBook(book *models.Book) error {
	result := r.db.Model(book).Where("id = ?", book.ID).Updates(book)
	if result.RowsAffected == 0 {
		return errors.New("Book not found")
	}
	return result.Error
}
