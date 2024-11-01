package service

import (
	"bookService/src/database/models"
	"bookService/src/repository"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
	return s.repo.CreateComment(comment)
}

func (s *CommentService) GetCommentsByBookID(bookID int) ([]models.Comment, error) {
	return s.repo.GetCommentsByBookID(bookID)
}

func (s *CommentService) GetCommentByID(id int) (*models.Comment, error) {
	return s.repo.GetCommentByID(id)
}
