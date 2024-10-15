package redis_repo

import (
	"bookService/src/database/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const (
	psychologyBooksKey        = "psychology_books"
	psychologyBooksKeyCounter = psychologyBooksKey + ":id"
)

func (r *redisRepository) AddPsychologyBook(ctx context.Context, book *models.Book) error {
	newID, err := r.redisClient.Incr(ctx, psychologyBooksKeyCounter).Result()
	if err != nil {
		return fmt.Errorf("could not increment ID counter: %w", err)
	}
	key := fmt.Sprintf("%s:%d", psychologyBooksKey, newID)
	book.ID = int(newID)

	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}

	return r.redisClient.Set(ctx, key, bookJSON, 0).Err()
}

func (r *redisRepository) GetPsychologyBooks(ctx context.Context) ([]*models.Book, error) {
	var books []*models.Book

	// Предположим, что у нас есть шаблон для ключей, и мы знаем максимальный ID
	// Например, мы можем использовать значение счетчика для определения максимального ID
	maxID, err := r.redisClient.Get(ctx, psychologyBooksKeyCounter).Int()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("could not get max ID from counter: %w", err)
	}

	for i := 1; i <= maxID; i++ {
		key := fmt.Sprintf("%s:%d", psychologyBooksKey, i)
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

func (r *redisRepository) GetPsychologyBook(ctx context.Context, id int) (*models.Book, error) {
	key := fmt.Sprintf("%s:%d", psychologyBooksKey, id)
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

func (r *redisRepository) DeletePsychologyBook(ctx context.Context, id int) error {
	key := fmt.Sprintf("%s:%d", psychologyBooksKey, id)
	return r.redisClient.Del(ctx, key).Err()
}
