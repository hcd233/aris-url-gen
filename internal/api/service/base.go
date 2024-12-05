package service

import "github.com/hcd233/Aris-url-gen/internal/api/dto"

type ShortURLService interface {
	GenerateShortURL(request *dto.GenerateShortURLRequest) (response *dto.GenerateShortURLResponse, err error)
	GetOriginalURL(request *dto.GetOriginalURLRequest) (response *dto.GetOriginalURLResponse, err error)
}
