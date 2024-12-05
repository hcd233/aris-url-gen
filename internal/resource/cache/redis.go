package cache

import (
	"context"
	"fmt"

	"github.com/hcd233/Aris-url-gen/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
)

var rdb *redis.Client

// GetRedisClient 获取Redis客户端
//
//	@return *redis.Client
//	@author centonhuang
//	@update 2024-12-05 16:16:07
func GetRedisClient() *redis.Client {
	return rdb
}

// InitCache 初始化Redis客户端
//
//	@author centonhuang
//	@update 2024-12-05 16:16:10
func InitCache() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       0,
	})

	_ = lo.Must1(rdb.Ping(context.Background()).Result())
}
