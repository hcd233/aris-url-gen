package util

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func GenerateRandomCode(originalURL string, length int) (randomCode string) {
	hasher := md5.New()
	hasher.Write([]byte(originalURL + time.Now().String())) // 加入时间戳增加随机性
	hash := hasher.Sum(nil)

	const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(base62Chars[hash[i]%62])
	}

	randomCode = builder.String()
	return
}

func ConstructFullShortURL(domainName string, path string, randomCode string) (fullShortURL string) {
	fullShortURL = fmt.Sprintf("http://%s/%s/%s", domainName, path, randomCode)
	return
}

func ProcessURL(originalURL string) (processedURL string) {
	parsedURL, err := url.Parse(originalURL)
	if err != nil {
		return
	}
	processedURL = fmt.Sprintf("http://%s/%s", parsedURL.Host, parsedURL.Path)
	return
}
