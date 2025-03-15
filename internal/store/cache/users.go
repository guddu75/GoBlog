package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/guddu75/goblog/internal/store"
)

type UserStore struct {
	client *redis.Client
}

const UserExpTime = time.Minute

func (s *UserStore) Get(ctx context.Context, key int64) (*store.User, error) {
	cacheKey := fmt.Sprintf("user-%v", key)

	data, err := s.client.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user store.User

	if data != "" {
		if err := json.Unmarshal([]byte(data), &user); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (s *UserStore) Set(ctx context.Context, user *store.User) error {

	cacheKey := fmt.Sprintf("user-%v", user.ID)

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return s.client.Set(ctx, cacheKey, data, UserExpTime).Err()
}

func (s *UserStore) Delete(ctx context.Context, key int64) error {
	cacheKey := fmt.Sprintf("user-%v", key)
	return s.client.Del(ctx, cacheKey).Err()
}
