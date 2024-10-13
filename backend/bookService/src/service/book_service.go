package service

import (
	"bookService/src/database/models"
	"bookService/src/repository"
	"log/slog"
)

type BookService struct {
	repo   repository.BookRepository
	Logger *slog.Logger
}

func NewBookService(repo repository.BookRepository, logger *slog.Logger) *BookService {
	return &BookService{repo: repo, Logger: logger}
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.repo.CreateBook(book)
}

func (s *BookService) GetBookByID(book *models.Book, bookID int) error {
	return s.repo.GetBookByID(book, bookID)

}

func (s *BookService) GetBooks() ([]models.Book, error) {
	return s.repo.GetBooks()
}

func (s *BookService) LikeBook(book *models.Book) error {
	return s.repo.LikeBook(book)
}

func (s *BookService) DislikeBook(book *models.Book) error {
	return s.repo.DislikeBook(book)
}

func (s *BookService) SearchByKeyword(keyword string) ([]models.Book, error) {
	return s.repo.SearchByKeyword(keyword)
}

func (s *BookService) DeleteBook(book *models.Book) error {
	return nil
}
