package redis_service

import (
	"bookService/src/database/models"
	"context"
)

// Функции для категории Romance
func (s *RedisService) AddRomanceBook(book *models.Book) error {
	return s.repo.AddRomanceBook(context.Background(), book)
}

func (s *RedisService) GetRomanceBooks() ([]*models.Book, error) {
	return s.repo.GetRomanceBooks(context.Background())
}

func (s *RedisService) GetRomanceBook(id int) (*models.Book, error) {
	return s.repo.GetRomanceBook(context.Background(), id)
}

func (s *RedisService) DeleteRomanceBook(id int) error {
	return s.repo.DeleteRomanceBook(context.Background(), id)
}
