package redis_service

import (
	"bookService/src/database/models"
	"context"
)

// Функции для категории Top Books
func (s *RedisService) AddToTopBooks(book *models.Book) error {
	return s.repo.AddToTopBooks(context.Background(), book)
}

func (s *RedisService) GetTopBooks() ([]*models.Book, error) {
	return s.repo.GetTopBooks(context.Background())
}

func (s *RedisService) DeleteTopBook(id int) error {
	return s.repo.DeleteTopBook(context.Background(), id)
}
