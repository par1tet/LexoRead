package service

import (
	"bookService/src/database/models"
	"bookService/src/repository"
	"context"
)

type RedisService struct {
	repo repository.RedisRepository
}

func NewRedisService(repo repository.RedisRepository) *RedisService {
	return &RedisService{repo: repo}
}

func (s *RedisService) AddBook(book *models.Book, key string) error {
	return s.repo.AddBook(context.Background(), book, key)
}

func (s *RedisService) GetBook(id int, key string) (*models.Book, error) {
	return s.repo.GetBook(context.Background(), id, key)
}

func (s *RedisService) DeleteBook(id int, key string) error {
	return s.repo.DeleteBook(context.Background(), id, key)
}

func (s *RedisService) UpdateBook(book *models.Book, key string, id string) error {
	return s.repo.UpdateBook(context.Background(), book, key, id)
}

func (s *RedisService) GetBooks(key string) ([]*models.Book, error) {
	return s.repo.GetBooks(context.Background(), key)
}
