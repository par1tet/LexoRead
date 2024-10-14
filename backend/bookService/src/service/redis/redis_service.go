package redis_service

import (
	"bookService/src/database/models"
	"bookService/src/repository/redis"
	"context"
)

type RedisService struct {
	repo redis_repo.RedisRepository
}

func NewRedisService(repo redis_repo.RedisRepository) *RedisService {
	return &RedisService{repo: repo}
}

func (s *RedisService) AddBook(book *models.Book) error {
	return s.repo.AddBook(context.Background(), book)
}

func (s *RedisService) GetBook(id int) (*models.Book, error) {
	return s.repo.GetBook(context.Background(), id)
}

func (s *RedisService) DeleteBook(id int) error {
	return s.repo.DeleteBook(context.Background(), id)
}

func (s *RedisService) UpdateBook(book *models.Book) error {
	return s.repo.UpdateBook(context.Background(), book)
}
