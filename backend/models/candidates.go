package models

import "gorm.io/gorm"

// Candidate 表示候选人的结构
type Candidate struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Address  string `gorm:"type:varchar(255);not null;unique" json:"address"`
	Votes    uint   `gorm:"default:0" json:"votes"`
	IsActive bool   `gorm:"default:true" json:"is_active"` // 标记候选人是否有效
}

// 自动迁移
func MigrateCandidates(db *gorm.DB) error {
	return db.AutoMigrate(&Candidate{})
}
