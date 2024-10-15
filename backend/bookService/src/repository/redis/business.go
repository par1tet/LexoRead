package redis_repo

import (
	"bookService/src/database/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const (
	businessBooksKey  = "Business_books"
	counterBusinessId = "Business_books:id"
)

func (r *redisRepository) AddBusinessBook(ctx context.Context, book *models.Book) error {
	newID, err := r.redisClient.Incr(ctx, counterBusinessId).Result()
	if err != nil {
		return fmt.Errorf("could not increment ID counter: %w", err)
	}
	key := fmt.Sprintf("%s:%d", businessBooksKey, newID)
	book.ID = int(newID)

	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}

	return r.redisClient.Set(ctx, key, bookJSON, 0).Err() // Сохраняем книгу как строку по уникальному ключу
}

func (r *redisRepository) GetBusinessBooks(ctx context.Context) ([]*models.Book, error) {
	var books []*models.Book

	// Предположим, что у нас есть шаблон для ключей, и мы знаем максимальный ID
	// Например, мы можем использовать значение счетчика для определения максимального ID
	maxID, err := r.redisClient.Get(ctx, counterBusinessId).Int()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("could not get max ID from counter: %w", err)
	}

	for i := 1; i <= maxID; i++ {
		key := fmt.Sprintf("%s:%d", businessBooksKey, i)
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

func (r *redisRepository) GetBusinessBook(ctx context.Context, id int) (*models.Book, error) {
	key := fmt.Sprintf("%s:%d", businessBooksKey, id) // Создаем ключ для получения книги
	bookJSON, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	var book models.Book
	if err := json.Unmarshal([]byte(bookJSON), &book); err != nil {
		return nil, fmt.Errorf("could not unmarshal book JSON: %w", err)
	}

	return &book, nil
}

func (r *redisRepository) DeleteBusinessBook(ctx context.Context, id int) error {
	key := fmt.Sprintf("%s:%d", businessBooksKey, id)
	return r.redisClient.Del(ctx, key).Err()
}
