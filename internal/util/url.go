// Package util 提供URL处理相关的工具函数
//
//	update 2024-12-07 01:02:11
package util

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/hcd233/Aris-url-gen/internal/config"
)

// GenerateRandomCode 生成随机码
//
//	param originalURL string
//	param length int
//	return randomCode string
//	author centonhuang
//	update 2024-12-05 16:16:44
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

// ConstructFullShortURL 构造完整的短URL
//
//	param domainName string
//	param path string
//	param randomCode string
//	return fullShortURL string
//	author centonhuang
//	update 2024-12-05 16:16:51
func ConstructFullShortURL(domainName string, path string, randomCode string) (fullShortURL string) {
	var scheme string
	if config.APIMode == config.ModeProd {
		scheme = "https"
	} else {
		scheme = "http"
	}
	fullShortURL = fmt.Sprintf("%s://%s/%s/%s", scheme, domainName, path, randomCode)
	return
}

// ProcessURL 处理URL
//
//	param originalURL string
//	return processedURL string
//	author centonhuang
//	update 2024-12-05 16:16:56
func ProcessURL(originalURL string) (processedURL string, err error) {
	parsedURL, err := url.Parse(originalURL)
	if err != nil {
		return
	}

	parsedURL.Scheme = "https"
	processedURL = parsedURL.String()
	return
}
