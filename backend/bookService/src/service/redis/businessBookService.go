package redis_service

import (
	"bookService/src/database/models"
	"context"
)

// Функции для категории Business
func (s *RedisService) AddBusinessBook(book *models.Book) error {
	return s.repo.AddBusinessBook(context.Background(), book)
}

func (s *RedisService) GetBusinessBooks() ([]*models.Book, error) {
	return s.repo.GetBusinessBooks(context.Background())
}

func (s *RedisService) GetBusinessBook(id int) (*models.Book, error) {
	return s.repo.GetBusinessBook(context.Background(), id)
}

func (s *RedisService) DeleteBusinessBook(id int) error {
	return s.repo.DeleteBusinessBook(context.Background(), id)
}
