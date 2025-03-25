package models

type Vote struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint `gorm:"not null"`
	CandidateID uint
}
