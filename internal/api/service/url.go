package service

import (
	"database/sql"
	"time"

	"github.com/hcd233/Aris-url-gen/internal/api/dto"
	"github.com/hcd233/Aris-url-gen/internal/config"
	"github.com/hcd233/Aris-url-gen/internal/resource/database"
	"github.com/hcd233/Aris-url-gen/internal/resource/database/dao"
	"github.com/hcd233/Aris-url-gen/internal/resource/database/model"
	"github.com/hcd233/Aris-url-gen/internal/util"
	"gorm.io/gorm"
)

type shortenUrlService struct {
	db     *gorm.DB
	urlDAO *dao.URLDAO
}

func NewShortenUrlService() ShortURLService {
	return &shortenUrlService{
		db:     database.GetDBInstance(),
		urlDAO: dao.GetURLDAO(),
	}
}

func (s *shortenUrlService) GenerateShortURL(request *dto.GenerateShortUrlRequest) (response *dto.GenerateShortURLResponse, err error) {
	url, err := s.urlDAO.GetByOriginalUrl(s.db, request.OriginalURL, []string{"id", "short_url"}, []string{})
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if url.ID != 0 {
		return &dto.GenerateShortURLResponse{
			ShortURL: util.ConstructFullShortURL(config.DomainName, "v1/shortURL", url.ShortURL),
		}, nil
	}

	shortURL := util.GenerateShortURL(request.OriginalURL, 8)

	url = &model.URL{
		OriginalURL: request.OriginalURL,
		ShortURL:    shortURL,
		ExpireAt:    sql.NullTime{Valid: false},
	}

	if request.ExpireDays > 0 {
		url.ExpireAt = sql.NullTime{Time: time.Now().AddDate(0, 0, int(request.ExpireDays)), Valid: true}
	}

	err = s.urlDAO.Create(s.db, url)
	if err != nil {
		return
	}

	return &dto.GenerateShortURLResponse{
		ShortURL: util.ConstructFullShortURL(config.DomainName, "v1/shortURL", shortURL),
	}, nil
}

func (s *shortenUrlService) GetOriginalURL(request *dto.GetOriginalUrlRequest) (response *dto.GetOriginalURLResponse, err error) {
	url, err := s.urlDAO.GetByShortUrl(s.db, request.ShortURL, []string{"original_url"}, []string{})
	if err != nil {
		return
	}

	return &dto.GetOriginalURLResponse{
		OriginalURL: url.OriginalURL,
	}, nil
}
