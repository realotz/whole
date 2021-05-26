package captcha

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	redis *redis.Client
}

func NewRedisStore(rdc *redis.Client) *RedisStore {
	return &RedisStore{
		redis: rdc,
	}
}

// Set sets the digits for the captcha id.
func (s *RedisStore) Set(id string, value string) {
	s.redis.Set(context.Background(), id, value, 0)
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (s *RedisStore) Get(id string, clear bool) string {
	if clear {
		defer func() {
			s.redis.Del(context.Background(), id)
		}()
	}
	return s.redis.Get(context.Background(), id).String()
}

//Verify captcha's answer directly
func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	val := s.Get(id, clear)
	return val == answer
}
