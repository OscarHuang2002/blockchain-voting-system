package controllers

import (
	"backend/models"
	"backend/services"
	"encoding/hex"
	"math/big"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// VerifyVote 验证投票记录
func VerifyVote(c *gin.Context) {
    address := c.Query("address")
    projectIdStr := c.Query("projectId")
    blockHashHex := c.Query("blockHash")

    if address == "" || projectIdStr == "" || blockHashHex == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "参数不完整"})
        return
    }

    // 将项目ID转为大整数
    projectId, ok := new(big.Int).SetString(projectIdStr, 10)
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的项目ID格式"})
        return
    }

    // 获取合约服务
    contractService := services.GetGlobalContractService()
    if contractService == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "合约服务未初始化"})
        return
    }

    // 检查用户是否在该项目中投票
    hasVoted, err := contractService.CheckUserVoted(projectId, address)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "检查投票状态失败: " + err.Error()})
        return
    }

    if !hasVoted {
        c.JSON(http.StatusBadRequest, gin.H{"error": "用户在此项目中未投票", "verified": false})
        return
    }

    // 获取投票详情
    voteDetail, err := contractService.GetVoteDetail(projectId, address)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取投票详情失败: " + err.Error()})
        return
    }

    // 将合约存储的区块哈希转为十六进制字符串
    contractBlockHashHex := "0x" + hex.EncodeToString(voteDetail.BlockHash[:])

    // 规范化输入的区块哈希格式（去除可能的0x前缀并转为小写）
    // 规范化输入的区块哈希格式（去除可能的0x前缀并转为小写）
    inputBlockHash := strings.ToLower(blockHashHex)
    inputBlockHash = strings.TrimPrefix(inputBlockHash, "0x")
    
    contractBlockHash := strings.ToLower(contractBlockHashHex)
    contractBlockHash = strings.TrimPrefix(contractBlockHash, "0x")

    // 验证区块哈希是否匹配
    verified := inputBlockHash == contractBlockHash

    // 如果验证成功，获取更多信息用于显示
    var result models.VoteVerificationResult
    result.Verified = verified

    if verified {
        // 获取项目信息
        project, err := contractService.GetProject(projectId)
        if err == nil {
            // 获取候选人信息
            candidate, err := contractService.GetCandidate(projectId, voteDetail.CandidateId)
            if err == nil {
                result.VoterAddress = address
                result.ProjectId = projectId.String()
                result.ProjectName = project.Description
                result.CandidateId = voteDetail.CandidateId.String()
                result.CandidateName = candidate.Name
                result.CandidateImage = candidate.ImageUrl
                result.BlockHash = contractBlockHashHex
            }
        }
    }

    c.JSON(http.StatusOK, result)
}