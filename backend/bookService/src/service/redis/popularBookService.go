package redis_service

import (
	"bookService/src/database/models"
	"context"
)

// Функции для категории Popular
func (s *RedisService) AddPopularBook(book *models.Book) error {
	return s.repo.AddPopularBook(context.Background(), book)
}

func (s *RedisService) GetPopularBooks() ([]*models.Book, error) {
	return s.repo.GetPopularBooks(context.Background())
}

func (s *RedisService) GetPopularBook(id int) (*models.Book, error) {
	return s.repo.GetPopularBook(context.Background(), id)
}

func (s *RedisService) DeletePopularBook(id int) error {
	return s.repo.DeletePopularBook(context.Background(), id)
}
