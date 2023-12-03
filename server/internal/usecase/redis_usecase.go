package usecase

import "time"

type RedisUsecase interface {
	Set(key string, value interface{}) error
	SetNx(key string, value interface{}, expireAt time.Duration) error
	Get(key string) (value interface{}, err error)
}
