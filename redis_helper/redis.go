package redishelper

import (
	"context"
	stderr "errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var (
	ErrOverMaxValue = stderr.New("OverMaxValue")
)

func NewClientByAssigned(host string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: host,
		DB:   db,
	})

	ctx := context.Background()
	err := client.Set(ctx, "test", "", 0).Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return client, nil
}

func IncWithMaxValue(ctx context.Context, c *redis.Client, key string, max int64, ttl time.Duration) error {
	cnt := c.Incr(ctx, key).Val()
	if cnt > max {
		return ErrOverMaxValue
	}

	c.Expire(ctx, key, ttl)

	return nil
}
