package redis_service

import (
	"bookService/src/database/models"
	"context"
)

func (s *RedisService) AddSelfDevelopmentBook(book *models.Book) error {
	return s.repo.AddSelfDevelopmentBook(context.Background(), book)
}

func (s *RedisService) GetSelfDevelopmentBooks() ([]*models.Book, error) {
	return s.repo.GetSelfDevelopmentBooks(context.Background())
}

func (s *RedisService) GetSelfDevelopmentBook(id int) (*models.Book, error) {
	return s.repo.GetSelfDevelopmentBook(context.Background(), id)
}

func (s *RedisService) DeleteSelfDevelopmentBook(id int) error {
	return s.repo.DeleteSelfDevelopmentBook(context.Background(), id)
}
