package service

import "github.com/hcd233/Aris-url-gen/internal/api/dto"

// ShortURLService 短URL服务
//
//	author centonhuang
//	update 2024-12-05 16:13:21
type ShortURLService interface {
	GenerateShortURL(request *dto.GenerateShortURLRequest) (response *dto.GenerateShortURLResponse, err error)
	GetOriginalURL(request *dto.GetOriginalURLRequest) (response *dto.GetOriginalURLResponse, err error)
}
