package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/guddu75/goblog/internal/store"
)

type Storage struct {
	Users interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
		Delete(context.Context, int64) error
	}
}

func NewRedisStorage(client *redis.Client) Storage {
	return Storage{
		Users: &UserStore{client},
	}
}
