package dbdao

import (
	"time"

	"github.com/hcd233/Aris-url-gen/internal/resource/database/model"
	"gorm.io/gorm"
)

// UserDAO 用户DAO
//
//	@author centonhuang
//	@update 2024-10-17 02:30:24
type URLDAO struct {
	baseDAO[model.URL]
}

// GetByEmail 通过邮箱获取用户
//
//	@receiver dao *UserDAO
//	@param db *gorm.DB
//	@param email string
//	@param fields []string``
//	@return user *model.User
//	@return err error
//	@author centonhuang
//	@update 2024-10-17 05:08:00
func (dao *URLDAO) GetByOriginalURL(db *gorm.DB, originalURL string, fields, preloads []string) (url *model.URL, err error) {
	sql := db.Select(fields)
	for _, preload := range preloads {
		sql = sql.Preload(preload)
	}
	err = sql.Where(model.URL{OriginalURL: originalURL}).First(&url).Error
	return
}

func (dao *URLDAO) GetByShortURL(db *gorm.DB, shortURL string, fields, preloads []string) (url *model.URL, err error) {
	sql := db.Select(fields)
	for _, preload := range preloads {
		sql = sql.Preload(preload)
	}
	err = sql.Where(model.URL{ShortURL: shortURL}).First(&url).Error
	return
}

func (dao *URLDAO) BatchGetExpiredURLs(db *gorm.DB, fields, preloads []string) (urls []*model.URL, err error) {
	sql := db.Select(fields)
	for _, preload := range preloads {
		sql = sql.Preload(preload)
	}
	err = sql.Where("expire_at < ?", time.Now()).Find(&urls).Error
	return
}

// BatchGetHotURLs 批量获取热门URL
func (dao *URLDAO) BatchGetHotURLs(db *gorm.DB, offset, limit int) ([]*model.URL, error) {
	var urls []*model.URL
	err := db.Model(&model.URL{}).
		Where("expire_at > ? OR expire_at IS NULL", time.Now()).
		Order("id DESC"). // 这里可以根据实际需求修改排序规则
		Offset(offset).
		Limit(limit).
		Find(&urls).Error
	return urls, err
}
