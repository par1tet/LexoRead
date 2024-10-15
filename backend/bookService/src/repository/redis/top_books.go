package redis_repo

import (
	"bookService/src/database/models"
	"context"
	"encoding/json"
	"fmt"
)

func (r *redisRepository) AddToTopBooks(ctx context.Context, book *models.Book) error {
	topBooksKey := "top_books"
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return r.redisClient.RPush(ctx, topBooksKey, bookJSON).Err()
}

func (r *redisRepository) GetTopBooks(ctx context.Context) ([]*models.Book, error) {
	topBooksJSON, err := r.redisClient.LRange(ctx, "top_books", 0, 99).Result()
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

func (r *redisRepository) DeleteTopBook(ctx context.Context, id int) error {
	key := fmt.Sprintf("top_books:%d", id)
	return r.redisClient.Del(ctx, key).Err()
}
