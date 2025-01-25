package dbdao

import (
	"sync"
)

var (
	urlDAOSingleton *URLDAO

	urlOnce sync.Once
)

// GetURLDAO 获取URLDAO
//
//	return *URLDAO
//	author centonhuang
//	update 2024-12-05 16:07:23
func GetURLDAO() *URLDAO {
	urlOnce.Do(func() {
		urlDAOSingleton = &URLDAO{}
	})
	return urlDAOSingleton
}
