package cachedao

import (
	"sync"

	"github.com/hcd233/Aris-url-gen/internal/resource/cache"
)

var (
	urlCacheDAOSingleton *urlCacheDAO

	urlOnce sync.Once
)

func GetURLCacheDAO() *urlCacheDAO {
	urlOnce.Do(func() {
		urlCacheDAOSingleton = &urlCacheDAO{
			rdb: cache.GetRedisClient(),
		}
	})
	return urlCacheDAOSingleton
}
