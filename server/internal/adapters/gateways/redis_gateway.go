package gateways

import (
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type Redis interface {
	Set(key string, value interface{}) error
	SetNx(key string, value interface{}, expireAt time.Duration) error
	Get(key string) (value interface{}, err error)
}

type RedisGateway struct {
	Redis Redis
}

func NewRedisGateway(redis Redis) usecase.RedisUsecase {
	return &RedisGateway{
		Redis: redis,
	}
}

func (r *RedisGateway) Set(key string, value interface{}) error {
	return r.Redis.Set(key, value)
}

func (r *RedisGateway) SetNx(key string, value interface{}, expireAt time.Duration) error {
	return r.Redis.SetNx(key, value, expireAt)
}

func (r *RedisGateway) Get(key string) (value interface{}, err error) {
	return r.Redis.Get(key)
}
