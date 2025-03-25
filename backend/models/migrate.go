package models

import (
	"gorm.io/gorm"
)

// Migrate 执行数据库迁移
func Migrate(db *gorm.DB) error {
	// 在这里添加你的模型
	db.AutoMigrate(&User{}, &Vote{})
	return db.AutoMigrate(&User{}, &Vote{})
}
