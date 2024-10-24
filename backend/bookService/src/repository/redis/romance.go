package redis_repo

import (
	"bookService/src/database/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const romanceBooksKey = "romance_books"

func (r *redisRepository) AddRomanceBook(ctx context.Context, book *models.Book) error {
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return r.redisClient.RPush(ctx, romanceBooksKey, bookJSON).Err()
}

func (r *redisRepository) GetRomanceBooks(ctx context.Context) ([]*models.Book, error) {
	topBooksJSON, err := r.redisClient.LRange(ctx, romanceBooksKey, 0, 99).Result()
	if err != nil {
		return nil, err
	}

	var books []*models.Book
	for _, bookJSON := range topBooksJSON {
		var book models.Book
		if err := json.Unmarshal([]byte(bookJSON), &book); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *redisRepository) GetRomanceBook(ctx context.Context, id int) (*models.Book, error) {
	key := fmt.Sprintf("%s:%d", romanceBooksKey, id)
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

func (r *redisRepository) DeleteRomanceBook(ctx context.Context, id int) error {
	key := fmt.Sprintf("%s:%d", romanceBooksKey, id)
	return r.redisClient.Del(ctx, key).Err()
}
