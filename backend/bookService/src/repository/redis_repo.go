package repository

import (
	"bookService/src/database/models"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	GetBook(ctx context.Context, id int, key_redis string) (*models.Book, error)
	AddBook(ctx context.Context, book *models.Book, key string) error
	DeleteBook(ctx context.Context, id int, key_redis string) error
	UpdateBook(ctx context.Context, book *models.Book, key_redis string, id string) error
	GetBooks(ctx context.Context, key string) ([]*models.Book, error)
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

func (r *redisRepository) AddBook(ctx context.Context, book *models.Book, key_redis string) error {
	newID, err := r.redisClient.Incr(ctx, key_redis+":id").Result()
	if err != nil {
		return fmt.Errorf("could not increment ID counter: %w", err)
	}
	key := fmt.Sprintf("%s:%d", key_redis, newID)
	book.ID = int(newID)

	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}

	// Логируем добавление книги
	log.Printf("Adding book with key %s: %+v\n", key, book)

	return r.redisClient.Set(ctx, key, bookJSON, 0).Err()
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

func (r *redisRepository) UpdateBook(ctx context.Context, book *models.Book, key_redis string, id string) error {
	key := fmt.Sprintf("%s:%s", key_redis, id)
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return r.redisClient.Set(ctx, key, bookJSON, 0).Err()
}

func (r *redisRepository) GetBooks(ctx context.Context, key string) ([]*models.Book, error) {
	var books []*models.Book

	// Предположим, что у нас есть шаблон для ключей, и мы знаем максимальный ID
	// Например, мы можем использовать значение счетчика для определения максимального ID
	maxID, err := r.redisClient.Get(ctx, key+":id").Int()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("could not get max ID from counter: %w", err)
	}

	for i := 1; i <= maxID; i++ {
		key := fmt.Sprintf("%s:%d", key, i)
		bookJSON, err := r.redisClient.Get(ctx, key).Result()
		if err != nil {
			if err != redis.Nil {
				return nil, fmt.Errorf("could not get book from Redis: %w", err)
			}
			// Книга не найдена, продолжаем со следующей
			continue
		}

		var book models.Book
		if err := json.Unmarshal([]byte(bookJSON), &book); err != nil {
			return nil, fmt.Errorf("could not unmarshal book JSON: %w", err)
		}
		books = append(books, &book)
	}

	return books, nil
}
