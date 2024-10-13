package redis_repo

import (
	"bookService/src/database/models"
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	AddBook(ctx context.Context, book *models.Book) error
	GetBook(ctx context.Context, id int) (*models.Book, error)
	GetTopBooks(ctx context.Context) ([]*models.Book, error)
	DeleteBook(ctx context.Context, id int) error
	UpdateBook(ctx context.Context, book *models.Book) error
	AddToTopBooks(ctx context.Context, book *models.Book) error
	DeleteTopBook(ctx context.Context, id int) error
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
