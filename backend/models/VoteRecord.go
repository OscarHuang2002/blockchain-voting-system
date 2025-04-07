package models

type VoteRecord struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserAddress string `json:"user_address"` // 改为字符串类型存储用户钱包地址
	ProjectID   uint   `json:"project_id"`
	CandidateID uint   `json:"candidate_id"`
	BlockHash   string `json:"block_hash"`
}