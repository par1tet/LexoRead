package redis_repo

import (
	"bookService/src/database/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const redisBooksKey = "book"

func (r *redisRepository) AddBook(ctx context.Context, book *models.Book) error {
	key := fmt.Sprintf("%s:%d", redisBooksKey, book.ID)
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return r.redisClient.Set(ctx, key, bookJSON, 0).Err()
}

func (r *redisRepository) GetBook(ctx context.Context, id int) (*models.Book, error) {
	key := fmt.Sprintf("%s:%d", redisBooksKey, id)
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

func (r *redisRepository) DeleteBook(ctx context.Context, id int) error {
	key := fmt.Sprintf("%s:%d", redisBooksKey, id)
	return r.redisClient.Del(ctx, key).Err()
}

func (r *redisRepository) UpdateBook(ctx context.Context, book *models.Book) error {
	key := fmt.Sprintf("%s:%d", redisBooksKey, book.ID)
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return r.redisClient.Set(ctx, key, bookJSON, 0).Err()
}
