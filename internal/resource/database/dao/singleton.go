package dao

import (
	"sync"
)

var (
	urlDAOSingleton *URLDAO

	urlOnce sync.Once
)

func GetURLDAO() *URLDAO {
	urlOnce.Do(func() {
		urlDAOSingleton = &URLDAO{}
	})
	return urlDAOSingleton
}
