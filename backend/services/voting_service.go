package services

import (
	"context"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var contractService *ContractService

// 定义未初始化错误
var ErrContractNotInitialized = errors.New("ContractService is not initialized")

// 初始化 ContractService
func InitContractService(service *ContractService) {
	contractService = service
}

// 投票
func CastVote(projectID uint, candidateID uint, userAddress string) (string, error) {
    if contractService == nil || contractService.contract == nil {
        return "", ErrContractNotInitialized
    }

    // 检查用户是否已注册
    isRegistered, err := contractService.contract.Voters(nil, common.HexToAddress(userAddress))
    if err != nil {
        return "", err
    }

    // 直接使用布尔值
    if !isRegistered {
        // 注册用户
        tx, err := contractService.contract.RegisterVoter(contractService.auth, common.HexToAddress(userAddress))
        if err != nil {
            return "", err
        }
        _, err = contractService.client.TransactionReceipt(context.Background(), tx.Hash())
        if err != nil {
            return "", err
        }
    }

    // 检查用户是否已经在此项目中投票
    hasVoted, err := contractService.contract.HasUserVoted(nil, big.NewInt(int64(projectID)), common.HexToAddress(userAddress))
    if err != nil {
        return "", err
    }

    if hasVoted {
        return "", errors.New("用户已经在此项目中投过票")
    }

    // 投票
    auth := contractService.auth
    tx, err := contractService.contract.Vote(auth, big.NewInt(int64(projectID)), big.NewInt(int64(candidateID)))
    if err != nil {
        return "", err
    }

    receipt, err := contractService.client.TransactionReceipt(context.Background(), tx.Hash())
    if err != nil {
        return "", err
    }

    // 获取区块哈希
    blockHash := receipt.BlockHash.Hex()

    // 将投票信息保存到数据库
    // 注意: 确保您有数据库模型来存储这些信息
    // config.DB.Create(&models.VoteRecord{
    //     UserID:      userAddress,
    //     ProjectID:   projectID,
    //     CandidateID: candidateID,
    //     BlockHash:   blockHash,
    // })

    return blockHash, nil
}

// 获取用户在特定项目中的投票详情
func GetUserVoteDetail(projectID uint, userAddress string) (map[string]interface{}, error) {
    // 使用 GlobalContractService 而不是 contractService
    if GlobalContractService == nil || GlobalContractService.contract == nil {
        return nil, ErrContractNotInitialized
    }

    // 添加日志帮助调试
    log.Printf("检查用户 %s 在项目 %d 中的投票状态", userAddress, projectID)

    // 检查用户是否在此项目中投票
    hasVoted, err := GlobalContractService.contract.HasUserVoted(nil, big.NewInt(int64(projectID)), common.HexToAddress(userAddress))
    if err != nil {
        log.Printf("检查投票状态错误: %v", err)
        return nil, err
    }

    if !hasVoted {
        log.Printf("用户未投票")
        // 不返回错误，而是返回空结果和特殊状态码
        return map[string]interface{}{"hasVoted": false}, nil
    }

    // 从voteDetails映射获取投票详情
    voteDetail, err := GlobalContractService.contract.VoteDetails(nil, big.NewInt(int64(projectID)), common.HexToAddress(userAddress))
    if err != nil {
        log.Printf("获取投票详情错误: %v", err)
        return nil, err
    }

    // 获取候选人信息
    id, name, imageUrl, description, voteCount, isActive, err := GlobalContractService.contract.GetCandidate(
        nil, 
        big.NewInt(int64(projectID)), 
        voteDetail.CandidateId,
    )
    if err != nil {
        log.Printf("获取候选人信息错误: %v", err)
        return nil, err
    }

    return map[string]interface{}{
        "hasVoted":      true,
        "candidateId":   id.Uint64(),
        "candidateName": name,
        "imageUrl":      imageUrl,
        "description":   description,
        "voteCount":     voteCount.Uint64(),
        "isActive":      isActive,
        "blockHash":     common.Bytes2Hex(voteDetail.BlockHash[:]),
    }, nil
}

// 获取所有候选人
func (cs *ContractService) GetAllCandidates(projectID uint) ([]map[string]interface{}, error) {
    if cs == nil || cs.contract == nil {
        return nil, ErrContractNotInitialized
    }

    candidates, err := cs.contract.GetAllCandidates(nil, big.NewInt(int64(projectID)))
    if err != nil {
        log.Printf("获取候选人列表失败: %v", err)
        return nil, err
    }

    var result []map[string]interface{}
    for _, candidate := range candidates {
        // 只添加活跃状态的候选人
        if candidate.IsActive {
            result = append(result, map[string]interface{}{
                "id":          candidate.Id.Uint64(),
                "name":        candidate.Name,
                "imageUrl":    candidate.ImageUrl,
                "description": candidate.CandidateDescription,
                "votes":       candidate.VoteCount.Uint64(),
                "isActive":    candidate.IsActive,
            })
        }
    }

    return result, nil
}