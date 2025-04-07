package models

import "math/big"

// UserVoteHistory 用户投票历史记录
type UserVoteHistory struct {
    ProjectId       *big.Int `json:"projectId"`
    ProjectName     string   `json:"projectName"`
    CandidateId     *big.Int `json:"candidateId"`
    CandidateName   string   `json:"candidateName"`
    CandidateImage  string   `json:"candidateImage"`
    BlockHash       string   `json:"blockHash"`     // 使用字符串存储区块哈希
    VoteTime        string   `json:"voteTime,omitempty"` // 如果有投票时间记录
}

// UserVoteResponse 用户投票历史响应
type UserVoteResponse struct {
    Address   string            `json:"address"`
    VoteCount int               `json:"voteCount"`
    Votes     []UserVoteHistory `json:"votes"`
}