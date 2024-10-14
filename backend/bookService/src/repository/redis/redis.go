package redis_repo

import (
	"bookService/src/database/models"
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	AddBook(ctx context.Context, book *models.Book) error
	GetBook(ctx context.Context, id int) (*models.Book, error)
	DeleteBook(ctx context.Context, id int) error
	UpdateBook(ctx context.Context, book *models.Book) error

	GetTopBooks(ctx context.Context) ([]*models.Book, error)
	AddToTopBooks(ctx context.Context, book *models.Book) error
	DeleteTopBook(ctx context.Context, id int) error

	AddBusinessBook(ctx context.Context, book *models.Book) error
	GetBusinessBooks(ctx context.Context) ([]*models.Book, error)
	GetBusinessBook(ctx context.Context, id int) (*models.Book, error)
	DeleteBusinessBook(ctx context.Context, id int) error

	AddFictionBook(ctx context.Context, book *models.Book) error
	GetFictionBooks(ctx context.Context) ([]*models.Book, error)
	GetFictionBook(ctx context.Context, id int) (*models.Book, error)
	DeleteFictionBook(ctx context.Context, id int) error

	AddGamesBook(ctx context.Context, book *models.Book) error
	GetGamesBooks(ctx context.Context) ([]*models.Book, error)
	GetGamesBook(ctx context.Context, id int) (*models.Book, error)
	DeleteGamesBook(ctx context.Context, id int) error

	AddLiteratureBook(ctx context.Context, book *models.Book) error
	GetLiteratureBooks(ctx context.Context) ([]*models.Book, error)
	GetLiteratureBook(ctx context.Context, id int) (*models.Book, error)
	DeleteLiteratureBook(ctx context.Context, id int) error

	AddNoveltyBook(ctx context.Context, book *models.Book) error
	GetNoveltyBooks(ctx context.Context) ([]*models.Book, error)
	GetNoveltyBook(ctx context.Context, id int) (*models.Book, error)
	DeleteNoveltyBook(ctx context.Context, id int) error

	AddPopularBook(ctx context.Context, book *models.Book) error
	GetPopularBooks(ctx context.Context) ([]*models.Book, error)
	GetPopularBook(ctx context.Context, id int) (*models.Book, error)
	DeletePopularBook(ctx context.Context, id int) error

	AddPsychologyBook(ctx context.Context, book *models.Book) error
	GetPsychologyBooks(ctx context.Context) ([]*models.Book, error)
	GetPsychologyBook(ctx context.Context, id int) (*models.Book, error)
	DeletePsychologyBook(ctx context.Context, id int) error

	AddRomanceBook(ctx context.Context, book *models.Book) error
	GetRomanceBooks(ctx context.Context) ([]*models.Book, error)
	GetRomanceBook(ctx context.Context, id int) (*models.Book, error)
	DeleteRomanceBook(ctx context.Context, id int) error

	AddSelfDevelopmentBook(ctx context.Context, book *models.Book) error
	GetSelfDevelopmentBooks(ctx context.Context) ([]*models.Book, error)
	GetSelfDevelopmentBook(ctx context.Context, id int) (*models.Book, error)
	DeleteSelfDevelopmentBook(ctx context.Context, id int) error
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
