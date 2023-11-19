package gateways

type Redis interface {
	Set(key string, value interface{}) error
	Get(key string) (value interface{}, err error)
}

type RedisGateway struct {
	Redis Redis
}

func NewRedisGateway() Redis {
	return &RedisGateway{}
}

func (r *RedisGateway) Set(key string, value interface{}) error {
	return r.Redis.Set(key, value)
}

func (r *RedisGateway) Get(key string) (value interface{}, err error) {
	return r.Redis.Get(key)
}
