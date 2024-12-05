package cachedao

import (
	"sync"

	"github.com/hcd233/Aris-url-gen/internal/resource/cache"
)

var (
	urlCacheDAOSingleton *urlCacheDAO

	urlOnce sync.Once
)

// GetURLCacheDAO 获取URL缓存DAO
//
//	@return *urlCacheDAO
//	@author centonhuang
//	@update 2024-12-05 16:06:05
func GetURLCacheDAO() *urlCacheDAO {
	urlOnce.Do(func() {
		urlCacheDAOSingleton = &urlCacheDAO{
			rdb: cache.GetRedisClient(),
		}
	})
	return urlCacheDAOSingleton
}
