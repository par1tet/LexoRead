package redis_repo

import (
	"bookService/src/database/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

const (
	fictionBooksKey   = "fiction_books"
	counterFictionKey = fictionBooksKey + ":id"
)

func (r *redisRepository) AddFictionBook(ctx context.Context, book *models.Book) error {
	newID, err := r.redisClient.Incr(ctx, counterFictionKey).Result()
	if err != nil {
		return fmt.Errorf("could not increment ID counter: %w", err)
	}
	key := fmt.Sprintf("%s:%d", fictionBooksKey, newID)
	book.ID = int(newID)

	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}

	// Логируем добавление книги
	log.Printf("Adding book with key %s: %+v\n", key, book)

	return r.redisClient.Set(ctx, key, bookJSON, 0).Err()
}

func (r *redisRepository) GetFictionBooks(ctx context.Context) ([]*models.Book, error) {
	var books []*models.Book

	// Предположим, что у нас есть шаблон для ключей, и мы знаем максимальный ID
	// Например, мы можем использовать значение счетчика для определения максимального ID
	maxID, err := r.redisClient.Get(ctx, counterGamesKey).Int()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("could not get max ID from counter: %w", err)
	}

	for i := 1; i <= maxID; i++ {
		key := fmt.Sprintf("%s:%d", gamesBooksKey, i)
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

func (r *redisRepository) GetFictionBook(ctx context.Context, id int) (*models.Book, error) {
	key := fmt.Sprintf("%s:%d", fictionBooksKey, id)
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

func (r *redisRepository) DeleteFictionBook(ctx context.Context, id int) error {
	key := fmt.Sprintf("%s:%d", fictionBooksKey, id)
	return r.redisClient.Del(ctx, key).Err()
}
