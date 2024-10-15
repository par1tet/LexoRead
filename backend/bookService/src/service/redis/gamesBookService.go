package redis_service

import (
	"bookService/src/database/models"
	"context"
)

// Функции для категории Games
func (s *RedisService) AddGamesBook(book *models.Book) error {
	return s.repo.AddGamesBook(context.Background(), book)
}

func (s *RedisService) GetGamesBooks() ([]*models.Book, error) {
	return s.repo.GetGamesBooks(context.Background())
}

func (s *RedisService) GetGamesBook(id int) (*models.Book, error) {
	return s.repo.GetGamesBook(context.Background(), id)
}

func (s *RedisService) DeleteGamesBook(id int) error {
	return s.repo.DeleteGamesBook(context.Background(), id)
}
