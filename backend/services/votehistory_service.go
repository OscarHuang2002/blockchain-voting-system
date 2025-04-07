package services

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// 添加结构体定义
type ProjectInfo struct {
    Id          *big.Int
    Description string
    VotingOpen  bool
}

type CandidateInfo struct {
    Id         *big.Int
    Name       string
    ImageUrl   string
    Desc       string
    VoteCount  *big.Int
    IsActive   bool
}

type VoteDetailInfo struct {
    CandidateId *big.Int
    BlockHash   [32]byte
}

// CheckUserVoted 检查用户是否在指定项目中投过票
func (s *ContractService) CheckUserVoted(projectId *big.Int, address string) (bool, error) {
    userAddress := common.HexToAddress(address)
    return s.contract.HasUserVoted(nil, projectId, userAddress)
}

// GetProject 获取项目信息
func (s *ContractService) GetProject(projectId *big.Int) (*ProjectInfo, error) {
    id, description, votingOpen, err := s.contract.GetProjectInfo(nil, projectId)
    if err != nil {
        return nil, err
    }
    
    return &ProjectInfo{
        Id:          id,
        Description: description,
        VotingOpen:  votingOpen,
    }, nil
}

// GetVoteDetail 获取用户在特定项目中的投票详情
func (s *ContractService) GetVoteDetail(projectId *big.Int, address string) (*VoteDetailInfo, error) {
    // 转换地址字符串为 ethereum 地址
    userAddress := common.HexToAddress(address)
    
    // 调用合约方法获取投票详情
    voteDetail, err := s.contract.VoteDetails(nil, projectId, userAddress)
    if err != nil {
        return nil, err
    }
    
    return &VoteDetailInfo{
        CandidateId: voteDetail.CandidateId,
        BlockHash:   voteDetail.BlockHash,
    }, nil
}

// GetCandidate 获取候选人信息
func (s *ContractService) GetCandidate(projectId *big.Int, candidateId *big.Int) (*CandidateInfo, error) {
    id, name, imageUrl, desc, voteCount, isActive, err := s.contract.GetCandidate(nil, projectId, candidateId)
    if err != nil {
        return nil, err
    }
    
    return &CandidateInfo{
        Id:        id,
        Name:      name,
        ImageUrl:  imageUrl,
        Desc:      desc,
        VoteCount: voteCount,
        IsActive:  isActive,
    }, nil
}