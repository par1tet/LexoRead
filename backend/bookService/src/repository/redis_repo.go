package repository

import (
	"bookService/src/database/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	GetBook(ctx context.Context, id int, key_redis string) (*models.Book, error)
	AddBook(ctx context.Context, book *models.Book, key string) error
	DeleteBook(ctx context.Context, id int, key_redis string) error
	UpdateBook(ctx context.Context, book *models.Book, key_redis string) error
}

type redisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(addr string) RedisRepository {
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &redisRepository{redisClient: redisClient}
}

func (r *redisRepository) AddBook(ctx context.Context, book *models.Book, key string) error {
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return r.redisClient.RPush(ctx, key, bookJSON).Err()
}

func (r *redisRepository) GetBook(ctx context.Context, id int, key_redis string) (*models.Book, error) {
	key := fmt.Sprintf("%s:%d", key_redis, id)
	bookJSON, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	var book models.Book
	if err := json.Unmarshal([]byte(bookJSON), &book); err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *redisRepository) DeleteBook(ctx context.Context, id int, key_redis string) error {
	key := fmt.Sprintf("%s:%d", key_redis, id)
	return r.redisClient.Del(ctx, key).Err()
}

func (r *redisRepository) UpdateBook(ctx context.Context, book *models.Book, key_redis string) error {
	key := fmt.Sprintf("%s:%d", key_redis, book.ID)
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return r.redisClient.Set(ctx, key, bookJSON, 0).Err()
}
