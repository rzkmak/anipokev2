package backend

import (
	"errors"
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	Client *redis.Client
	TTL    time.Duration
}

func (r Redis) Get(key string) (int, error) {
	raw, err := r.Client.Get(key).Int()
	if err != nil {
		return 0, err
	}

	return raw, nil
}

func (r Redis) Set(key string) error {
	if key == "" {
		return errors.New("redis: set key should not be nil")
	}

	_, err := r.Client.Set(key, 1, r.TTL).Result()
	if err != nil {
		return err
	}
	return nil
}
