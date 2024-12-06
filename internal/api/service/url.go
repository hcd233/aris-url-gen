package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	cachedao "github.com/hcd233/Aris-url-gen/internal/api/dao/cache"
	dbdao "github.com/hcd233/Aris-url-gen/internal/api/dao/db"
	"github.com/hcd233/Aris-url-gen/internal/api/dto"
	"github.com/hcd233/Aris-url-gen/internal/config"
	"github.com/hcd233/Aris-url-gen/internal/logger"
	"github.com/hcd233/Aris-url-gen/internal/resource/database"
	"github.com/hcd233/Aris-url-gen/internal/resource/database/model"
	"github.com/hcd233/Aris-url-gen/internal/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	randomCodeLength = 8
	regenerateTimes  = 3

	nilCacheValue = "nil"

	cacheExpiration = 1 * time.Hour
)

var (
	errShortURLNotFound = errors.New("short url not found")
	errShortURLExpired  = errors.New("short url has expired")
)

type shortURLService struct {
	db          *gorm.DB
	urlDAO      *dbdao.URLDAO
	urlCacheDAO cachedao.URLCacheDAO
}

// NewShortURLService 创建短URL服务
//
//	@return ShortURLService
//	@author centonhuang
//	@update 2024-12-05 16:13:36
func NewShortURLService() ShortURLService {
	return &shortURLService{
		db:          database.GetDBInstance(),
		urlDAO:      dbdao.GetURLDAO(),
		urlCacheDAO: cachedao.GetURLCacheDAO(),
	}
}

// GenerateShortURL 生成短URL
//
//	@receiver s *shortURLService
//	@param request *dto.GenerateShortURLRequest
//	@return response *dto.GenerateShortURLResponse
//	@return err error
//	@author centonhuang
//	@update 2024-12-05 16:13:40
func (s *shortURLService) GenerateShortURL(request *dto.GenerateShortURLRequest) (response *dto.GenerateShortURLResponse, err error) {
	// 先查询缓存
	ctx := context.Background()

	originalURL, err := util.ProcessURL(request.OriginalURL)
	if err != nil {
		return
	}
	shortURL, err := s.urlCacheDAO.GetURLByOriginal(ctx, originalURL)
	if shortURL != "" && err == nil {
		return &dto.GenerateShortURLResponse{
			ShortURL: util.ConstructFullShortURL(config.DomainName, "v1/s", shortURL),
		}, nil
	}

	url, err := s.urlDAO.GetByOriginalURL(s.db, originalURL, []string{"id", "short_url"}, []string{})
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if url.ID != 0 {
		if err = s.urlCacheDAO.SetBidirectionalCache(ctx, url.ShortURL, originalURL, cacheExpiration); err != nil {
			logger.Logger.Error("[Cache] failed to set bidirectional cache", zap.Error(err))
		}
		logger.Logger.Info("[Cache] short url already exists", zap.String("short_url", url.ShortURL))
		return &dto.GenerateShortURLResponse{
			ShortURL: util.ConstructFullShortURL(config.DomainName, "v1/s", url.ShortURL),
		}, nil
	}

	var shortURLSuffix string
	for i := 0; i < regenerateTimes; i++ {
		shortURLSuffix = util.GenerateRandomCode(request.OriginalURL, randomCodeLength)

		url, err = s.urlDAO.GetByShortURL(s.db, shortURLSuffix, []string{"id"}, []string{})
		if err != nil && err != gorm.ErrRecordNotFound {
			return
		}
		if url.ID == 0 {
			break
		}
	}

	url = &model.URL{
		OriginalURL: originalURL,
		ShortURL:    shortURLSuffix,
		ExpireAt:    sql.NullTime{Valid: false},
	}

	if request.ExpireDays > 0 {
		url.ExpireAt = sql.NullTime{Time: time.Now().AddDate(0, 0, int(request.ExpireDays)), Valid: true}
	}

	err = s.urlDAO.Create(s.db, url)
	if err != nil {
		return
	}

	// 修改缓存写入部分
	var expiration time.Duration
	if url.ExpireAt.Valid {
		expiration = time.Until(url.ExpireAt.Time)
	} else {
		expiration = 24 * time.Hour
	}

	// 使用双向缓存
	if err = s.urlCacheDAO.SetBidirectionalCache(ctx, shortURLSuffix, url.OriginalURL, expiration); err != nil {
		logger.Logger.Error("[Cache] failed to set bidirectional cache", zap.Error(err))
	}

	return &dto.GenerateShortURLResponse{
		ShortURL: util.ConstructFullShortURL(config.DomainName, "v1/s", shortURLSuffix),
	}, nil
}

// GetOriginalURL 获取原始URL
//
//	@receiver s *shortURLService
//	@param request *dto.GetOriginalURLRequest
//	@return response *dto.GetOriginalURLResponse
//	@return err error
//	@author centonhuang
//	@update 2024-12-05 16:13:53
func (s *shortURLService) GetOriginalURL(request *dto.GetOriginalURLRequest) (response *dto.GetOriginalURLResponse, err error) {
	ctx := context.Background()

	// 先从缓存中查询
	originalURL, err := s.urlCacheDAO.GetURLByShort(ctx, request.ShortURL)
	if originalURL != "" && err == nil {
		// 检查是否是空缓存标记
		if originalURL == nilCacheValue {
			return nil, errShortURLNotFound
		}
		return &dto.GetOriginalURLResponse{
			OriginalURL: originalURL,
		}, nil
	}

	url, err := s.urlDAO.GetByShortURL(s.db, request.ShortURL, []string{"original_url", "expire_at"}, []string{})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			if err = s.urlCacheDAO.SetNilCacheByShortURL(ctx, request.ShortURL, cacheExpiration); err != nil {
				logger.Logger.Error("[Cache] failed to set nil cache", zap.Error(err))
			}
			logger.Logger.Warn("[Cache] short url not found", zap.String("short_url", request.ShortURL))
			return nil, errShortURLNotFound
		}
	}

	if url.ExpireAt.Valid && url.ExpireAt.Time.Before(time.Now()) {
		if err = s.urlCacheDAO.SetNilCacheByShortURL(ctx, request.ShortURL, cacheExpiration); err != nil {
			logger.Logger.Error("[Cache] failed to set nil cache", zap.Error(err))
		}
		logger.Logger.Warn("[Cache] short url has expired", zap.String("short_url", request.ShortURL))
		return nil, errShortURLExpired
	}

	// 使用双向缓存
	if err = s.urlCacheDAO.SetBidirectionalCache(ctx, request.ShortURL, url.OriginalURL, cacheExpiration); err != nil {
		logger.Logger.Error("[Cache]failed to set bidirectional cache", zap.Error(err))
	}

	return &dto.GetOriginalURLResponse{
		OriginalURL: url.OriginalURL,
	}, nil
}
