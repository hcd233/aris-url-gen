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

// Cron 定时任务接口
//
//	author centonhuang
//	update 2024-12-05 16:14:29
type Cron interface {
	Start()
}

// CleanExpiredURLsCron 清理过期的URL
//
//	author centonhuang
//	update 2024-12-05 16:14:35
type CleanExpiredURLsCron struct {
	cron   *cron.Cron
	db     *gorm.DB
	urlDAO *dbdao.URLDAO
}

// NewCleanExpiredURLsCron 创建清理过期的URL定时任务
//
//	return Cron
//	author centonhuang
//	update 2024-12-05 16:14:39
func NewCleanExpiredURLsCron() Cron {
	return &CleanExpiredURLsCron{
		cron: cron.New(
			cron.WithLogger(newCronLoggerAdapter("CleanExpiredURLsCron", logger.Logger)),
		),
		db:     database.GetDBInstance(),
		urlDAO: dbdao.GetURLDAO(),
	}
}

// Start 启动定时任务
//
//	receiver c *CleanExpiredURLsCron
//	author centonhuang
//	update 2024-12-05 16:14:45
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
