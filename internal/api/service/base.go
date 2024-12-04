package service

import "github.com/hcd233/Aris-url-gen/internal/api/dto"

type ShortURLService interface {
	GenerateShortURL(request *dto.GenerateShortUrlRequest) (response *dto.GenerateShortURLResponse, err error)
	GetOriginalURL(request *dto.GetOriginalUrlRequest) (response *dto.GetOriginalURLResponse, err error)
}
