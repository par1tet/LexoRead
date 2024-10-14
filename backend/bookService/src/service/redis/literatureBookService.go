package redis_service

import (
	"bookService/src/database/models"
	"context"
)

// Функции для категории Literature
func (s *RedisService) AddLiteratureBook(book *models.Book) error {
	return s.repo.AddLiteratureBook(context.Background(), book)
}

func (s *RedisService) GetLiteratureBooks() ([]*models.Book, error) {
	return s.repo.GetLiteratureBooks(context.Background())
}

func (s *RedisService) GetLiteratureBook(id int) (*models.Book, error) {
	return s.repo.GetLiteratureBook(context.Background(), id)
}

func (s *RedisService) DeleteLiteratureBook(id int) error {
	return s.repo.DeleteLiteratureBook(context.Background(), id)
}
