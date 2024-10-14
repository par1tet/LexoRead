package redis_service

import (
	"bookService/src/database/models"
	"context"
)

// Функции для категории Psychology
func (s *RedisService) AddPsychologyBook(book *models.Book) error {
	return s.repo.AddPsychologyBook(context.Background(), book)
}

func (s *RedisService) GetPsychologyBooks() ([]*models.Book, error) {
	return s.repo.GetPsychologyBooks(context.Background())
}

func (s *RedisService) GetPsychologyBook(id int) (*models.Book, error) {
	return s.repo.GetPsychologyBook(context.Background(), id)
}

func (s *RedisService) DeletePsychologyBook(id int) error {
	return s.repo.DeletePsychologyBook(context.Background(), id)
}
