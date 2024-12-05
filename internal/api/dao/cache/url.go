package cachedao

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type URLCacheDAO interface {
	GetURLByShort(ctx context.Context, shortURL string) (string, error)
	GetURLByOriginal(ctx context.Context, originalURL string) (string, error)
	SetBidirectionalCache(ctx context.Context, shortURL, originalURL string, expiration time.Duration) error
	SetNilCacheByShortURL(ctx context.Context, shortURL string, expiration time.Duration) error
}

const (
	originalURLKey = "originalURL:%s"
	shortURLKey    = "shortURL:%s"
)

type urlCacheDAO struct {
	rdb *redis.Client
}

func (dao *urlCacheDAO) GetURLByShort(ctx context.Context, shortURL string) (string, error) {
	cacheKey := fmt.Sprintf(shortURLKey, shortURL)
	return dao.rdb.Get(ctx, cacheKey).Result()
}
func (dao *urlCacheDAO) GetURLByOriginal(ctx context.Context, originalURL string) (string, error) {
	cacheKey := fmt.Sprintf(originalURLKey, originalURL)
	return dao.rdb.Get(ctx, cacheKey).Result()
}

func (dao *urlCacheDAO) SetBidirectionalCache(ctx context.Context, shortURL, originalURL string, expiration time.Duration) error {
	pipe := dao.rdb.Pipeline()

	// 设置 short->original 映射
	shortKey := fmt.Sprintf(shortURLKey, shortURL)
	pipe.Set(ctx, shortKey, originalURL, expiration)

	// 设置 original->short 映射
	originalKey := fmt.Sprintf(originalURLKey, originalURL)
	pipe.Set(ctx, originalKey, shortURL, expiration)

	_, err := pipe.Exec(ctx)
	return err
}

func (dao *urlCacheDAO) SetNilCacheByShortURL(ctx context.Context, shortURL string, expiration time.Duration) error {
	cacheKey := fmt.Sprintf(shortURLKey, shortURL)
	return dao.rdb.Set(ctx, cacheKey, "nil", expiration).Err()
}
