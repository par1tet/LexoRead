package redis_service

import (
	"bookService/src/database/models"
	"context"
)

// Функции для категории Fiction
func (s *RedisService) AddFictionBook(book *models.Book) error {
	return s.repo.AddFictionBook(context.Background(), book)
}

func (s *RedisService) GetFictionBooks() ([]*models.Book, error) {
	return s.repo.GetFictionBooks(context.Background())
}

func (s *RedisService) GetFictionBook(id int) (*models.Book, error) {
	return s.repo.GetFictionBook(context.Background(), id)
}

func (s *RedisService) DeleteFictionBook(id int) error {
	return s.repo.DeleteFictionBook(context.Background(), id)
}
