package infrastructure

import (
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/gateways"
)

type Redis struct {
	RDB *redis.Client
}

func NewRedis() gateways.Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return &Redis{
		RDB: rdb,
	}
}

func (r *Redis) Set(key string, value interface{}) error {
	err := r.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) SetNx(key string, value interface{}, expireAt time.Duration) error {
	err := r.SetNx(key, value, expireAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(key string) (value interface{}, err error) {
	value, err = r.Get(key)
	if err == redis.Nil {
		return value, fmt.Errorf("key does not exist: %v", key)
	}
	if err != nil {
		return value, err
	}
	return value, err
}
