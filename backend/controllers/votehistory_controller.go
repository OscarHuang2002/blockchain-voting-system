package controllers

import (
	"backend/models"
	"backend/services"
	"encoding/hex"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserVoteDetails 获取用户的所有投票历史
func GetUserVoteDetails(c *gin.Context) {
    // 从请求中获取钱包地址
    address := c.Query("address")
    if address == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "地址参数缺失"})
        return
    }

    // 获取合约服务
    contractService := services.GetGlobalContractService()
    if contractService == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "合约服务未初始化"})
        return
    }

    // 获取项目总数
    projectCount, err := contractService.GetProjectCount()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取项目总数失败: " + err.Error()})
        return
    }

    // 准备响应数据
    response := models.UserVoteResponse{
        Address:   address,
        VoteCount: 0,
        Votes:     []models.UserVoteHistory{},
    }

    // 遍历所有项目，查找用户的投票记录 (修改为使用 uint 类型)
    for i := uint(1); i <= projectCount; i++ {
        projectId := big.NewInt(int64(i)) // 将 uint 转为 *big.Int
        
        // 检查用户是否在该项目中投过票
        hasVoted, err := contractService.CheckUserVoted(projectId, address)
        if err != nil {
            continue // 如果查询出错，跳过此项目
        }

        // 如果用户在此项目中投过票，获取投票详情
        if hasVoted {
            // 获取项目信息
            project, err := contractService.GetProject(projectId)
            if err != nil {
                continue
            }

            // 获取用户的投票详情
            voteDetail, err := contractService.GetVoteDetail(projectId, address)
            if err != nil {
                continue
            }

            // 获取候选人信息
            candidateInfo, err := contractService.GetCandidate(projectId, voteDetail.CandidateId)
            if err != nil {
                continue
            }

            // 将[32]byte转换为十六进制字符串
            blockHashHex := "0x" + hex.EncodeToString(voteDetail.BlockHash[:])

            // 添加到响应中
            vote := models.UserVoteHistory{
                ProjectId:      projectId,
                ProjectName:    project.Description,
                CandidateId:    voteDetail.CandidateId,
                CandidateName:  candidateInfo.Name,
                CandidateImage: candidateInfo.ImageUrl,
                BlockHash:      blockHashHex,
            }
            response.Votes = append(response.Votes, vote)
            response.VoteCount++
        }
    }

    c.JSON(http.StatusOK, response)
}