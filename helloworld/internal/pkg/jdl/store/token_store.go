package store

import (
	"context"
	"encoding/json"
	"fmt"
	"helloworld/internal/pkg/jdl/auth"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ TokenStore = (*RedisTokenStore)(nil)

type TokenStore interface {
	GetToken(ctx context.Context) (*auth.Token, error)
	SetToken(ctx context.Context, token *auth.Token) error
}

type RedisTokenStore struct {
	client *redis.Client
	key    string
}

func NewRedisTokenStore(client *redis.Client, key string) *RedisTokenStore {
	return &RedisTokenStore{
		client: client,
		key:    key,
	}
}

func (r *RedisTokenStore) GetToken(ctx context.Context) (*auth.Token, error) {
	val, err := r.client.Get(ctx, r.key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var token auth.Token
	if err := json.Unmarshal([]byte(val), &token); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token from redis (key=%s, value=%q): %w", r.key, val, err)
	}

	return &token, nil
}

func (r *RedisTokenStore) SetToken(ctx context.Context, token *auth.Token) error {
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}

	ttl := time.Until(token.AccessExpire)
	if ttl <= 0 {
		return nil // refresh_token 已过期，不保存
	}

	if err := r.client.Set(ctx, r.key, data, ttl).Err(); err != nil {
		return fmt.Errorf("failed to set token in redis (key=%s): %w", r.key, err)
	}

	return nil
}
