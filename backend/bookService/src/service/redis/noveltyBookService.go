package redis_service

import (
	"bookService/src/database/models"
	"context"
)

// Функции для категории Novelty
func (s *RedisService) AddNoveltyBook(book *models.Book) error {
	return s.repo.AddNoveltyBook(context.Background(), book)
}

func (s *RedisService) GetNoveltyBooks() ([]*models.Book, error) {
	return s.repo.GetNoveltyBooks(context.Background())
}

func (s *RedisService) GetNoveltyBook(id int) (*models.Book, error) {
	return s.repo.GetNoveltyBook(context.Background(), id)
}

func (s *RedisService) DeleteNoveltyBook(id int) error {
	return s.repo.DeleteNoveltyBook(context.Background(), id)
}
