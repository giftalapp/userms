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

	log.Println("[INFO] Connected to Redis")

	return &Bucket{
		ctx: context.Background(),
		rc:  rc,
		exp: expiryTime,
	}, nil
}

func (b *Bucket) Get(key string) error {
	status := b.rc.Get(b.ctx, key)

	return status.Err()
}

func (b *Bucket) Set(key string, value string) error {
	status := b.rc.Set(b.ctx, key, value, b.exp)

	return status.Err()
}
