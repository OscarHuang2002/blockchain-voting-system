package models

// VoteVerificationResult 投票验证结果
type VoteVerificationResult struct {
    Verified       bool   `json:"verified"`
    VoterAddress   string `json:"voterAddress,omitempty"`
    ProjectId      string `json:"projectId,omitempty"`
    ProjectName    string `json:"projectName,omitempty"`
    CandidateId    string `json:"candidateId,omitempty"`
    CandidateName  string `json:"candidateName,omitempty"`
    CandidateImage string `json:"candidateImage,omitempty"`
    BlockHash      string `json:"blockHash,omitempty"`
}