package bucket

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Bucket struct {
	ctx context.Context
	rc  *redis.Client
	exp time.Duration
}

func NewBucket(redisURL string, expiryTime time.Duration) (*Bucket, error) {
	opt, err := redis.ParseURL(redisURL)

	if err != nil {
		return nil, err
	}

	rc := redis.NewClient(opt)

	if err := rc.Set(context.Background(), "test", "test", time.Second).Err(); err != nil {
		return nil, err
	}

	rc.FlushDB(context.Background())

	log.Println("[INFO] Connected to Redis")

	return &Bucket{
		ctx: context.Background(),
		rc:  rc,
		exp: expiryTime,
	}, nil
}

func (b *Bucket) Get(key string) (interface{}, time.Duration, error) {
	status := b.rc.Get(b.ctx, key)
	ttl, err := b.rc.TTL(b.ctx, key).Result()

	if err != nil {
		return nil, time.Second, err
	}

	return status.Val(), ttl, status.Err()
}

func (b *Bucket) Set(key string, value interface{}, exp time.Duration) error {
	if exp == time.Duration(0) {
		exp = b.exp
	}

	status := b.rc.Set(b.ctx, key, value, exp)

	return status.Err()
}

func (b *Bucket) Del(key string) {
	b.rc.Del(b.ctx, key)
}
