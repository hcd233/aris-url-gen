package util

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
)

func GenerateShortURL(originalURL string, length int) (shortURL string) {
	hasher := md5.New()
	hasher.Write([]byte(originalURL + time.Now().String())) // 加入时间戳增加随机性
	hash := hasher.Sum(nil)

	const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(base62Chars[hash[i]%62])
	}

	shortURL = builder.String()
	return
}

func ConstructFullShortURL(domainName string, path string, shortURL string) (fullShortURL string) {
	fullShortURL = fmt.Sprintf("http://%s/%s/%s", domainName, path, shortURL)
	return
}
