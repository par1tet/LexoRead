package repository

import (
	"bookService/src/database/models" // Убедитесь, что путь правильный
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9" // Импортируем нужную библиотеку
)

// Интерфейс для работы с репозиторием книг
type RedisRepository interface {
	AddBook(ctx context.Context, book *models.Book) error
	GetBook(ctx context.Context, id int) (*models.Book, error)
	GetTopBooks(ctx context.Context) ([]*models.Book, error)
	DeleteBook(ctx context.Context, id int) error
	UpdateBook(ctx context.Context, book *models.Book) error
	AddToTopBooks(ctx context.Context, book *models.Book) error
}

// Реализация интерфейса RedisRepository
type redisRepository struct {
	redisClient *redis.Client
}

// Конструктор для redisRepository
func NewRedisRepository(addr string) RedisRepository {
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr, // Укажите правильный адрес
	})
	return &redisRepository{redisClient: redisClient}
}

// Метод добавления книги
func (r *redisRepository) AddBook(ctx context.Context, book *models.Book) error {
	key := fmt.Sprintf("book:%d", book.ID)
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return r.redisClient.Set(ctx, key, bookJSON, 0).Err()
}

// Метод получения книги по ID
func (r *redisRepository) GetBook(ctx context.Context, id int) (*models.Book, error) {
	key := fmt.Sprintf("book:%d", id)
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

// Метод получения топ-100 книг
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

// Метод удаления книги
func (r *redisRepository) DeleteBook(ctx context.Context, id int) error {
	key := fmt.Sprintf("book:%d", id)
	return r.redisClient.Del(ctx, key).Err()
}

// Метод обновления книги
func (r *redisRepository) UpdateBook(ctx context.Context, book *models.Book) error {
	key := fmt.Sprintf("book:%d", book.ID)
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return r.redisClient.Set(ctx, key, bookJSON, 0).Err()
}

// Метод для добавления книги в топ-100
func (r *redisRepository) AddToTopBooks(ctx context.Context, book *models.Book) error {
	topBooksKey := "top_books"
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return r.redisClient.RPush(ctx, topBooksKey, bookJSON).Err()
}
