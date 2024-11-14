package repository

import (
	"bookService/src/database/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	GetCommentsByBookID(bookID int) ([]models.Comment, error)
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepo{db: db}
}

func (r *commentRepo) CreateComment(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepo) GetCommentsByBookID(bookID int) ([]models.Comment, error){
	var comments []models.Comment
	err := r.db.Where("book_id = ?", bookID).Find(&comments).Error
	return comments, err
}
