// Package model defines the database schema for the model.
//
//	@update 2024-06-22 09:33:43
package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	ID          uint         `json:"id" gorm:"column:id;primary_key;auto_increment;comment:URL ID"`
	OriginalURL string       `json:"original_url" gorm:"column:original_url;not null;index:idx_original_url;comment:原始URL"`
	ShortURL    string       `json:"short_url" gorm:"column:short_url;not null;unique;comment:短URL"`
	ExpireAt    sql.NullTime `json:"expire_at" gorm:"column:expire_at;not null;index:idx_expire_at;comment:过期时间"`
}
