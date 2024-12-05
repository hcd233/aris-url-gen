package cron

import (
	dbdao "github.com/hcd233/Aris-url-gen/internal/api/dao/db"
	"github.com/hcd233/Aris-url-gen/internal/logger"
	"github.com/hcd233/Aris-url-gen/internal/resource/database"
	"github.com/hcd233/Aris-url-gen/internal/resource/database/model"
	"github.com/robfig/cron/v3"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Cron interface {
	Start()
}

type CleanExpiredURLsCron struct {
	cron   *cron.Cron
	db     *gorm.DB
	urlDAO *dbdao.URLDAO
}

func NewCleanExpiredURLsCron() Cron {
	return &CleanExpiredURLsCron{
		cron: cron.New(
			cron.WithLogger(newCronLoggerAdapter("CleanExpiredURLsCron", logger.Logger)),
		),
		db:     database.GetDBInstance(),
		urlDAO: dbdao.GetURLDAO(),
	}
}

func (c *CleanExpiredURLsCron) Start() {
	// debug set 10 seconds
	// c.cron.AddFunc("@every 10s", c.cleanExpiredURLs)
	c.cron.AddFunc("@daily", c.cleanExpiredURLs)
	c.cron.Start()
}

func (c *CleanExpiredURLsCron) cleanExpiredURLs() {
	urls, err := c.urlDAO.BatchGetExpiredURLs(c.db, []string{"id"}, []string{})
	if err != nil {
		logger.Logger.Error("[CleanExpiredURLsCron] cleanExpiredURLs error", zap.Error(err))
		return
	}

	logger.Logger.Info("[CleanExpiredURLsCron] cleanExpiredURLs stats", zap.Int("count", len(urls)))

	expiredURLs := lo.Map(urls, func(url *model.URL, idx int) model.URL {
		return *url
	})

	err = c.urlDAO.BatchDelete(c.db, &expiredURLs)
	if err != nil {
		logger.Logger.Error("[CleanExpiredURLsCron] cleanExpiredURLs error", zap.Error(err))
		return
	}

	logger.Logger.Info("[CleanExpiredURLsCron] cleanExpiredURLs success")
}
