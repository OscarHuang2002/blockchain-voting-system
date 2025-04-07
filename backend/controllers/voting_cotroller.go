package controllers

import (
	"backend/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InitContractService 初始化合约服务
func InitContractService(service *services.ContractService) {
    contractService = service
}

// 投票
func CastVote(c *gin.Context) {
    var voteData struct {
        ProjectID   uint   `json:"project_id" binding:"required"`
        CandidateID uint   `json:"candidate_id" binding:"required"`
        UserAddress string `json:"user_address" binding:"required"`
    }

    if err := c.ShouldBindJSON(&voteData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
        return
    }

    blockHash, err := services.CastVote(voteData.ProjectID, voteData.CandidateID, voteData.UserAddress)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "投票失败: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "投票成功",
        "blockHash": blockHash,
    })
}

// 获取用户在特定项目中的投票详情
func GetUserVoteDetail(c *gin.Context) {
    projectIDStr := c.Query("project_id")
    userAddress := c.Query("user_address")

    if projectIDStr == "" || userAddress == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "缺少项目ID或用户地址"})
        return
    }

    projectID, err := strconv.Atoi(projectIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的项目ID"})
        return
    }

    voteDetail, err := services.GetUserVoteDetail(uint(projectID), userAddress)
    if err != nil {
        // 这里处理真正的错误
        log.Printf("获取投票详情失败: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取投票详情失败", "details": err.Error()})
        return
    }

    // 检查是否有投票记录
    if voteDetail["hasVoted"] == false {
        // 用户未投票，返回200状态码但标记未投票
        c.JSON(http.StatusOK, gin.H{"hasVoted": false})
        return
    }

    // 用户已投票，返回详细信息
    c.JSON(http.StatusOK, gin.H{"voteDetail": voteDetail, "hasVoted": true})
}

// 获取所有候选人
func GetAllCandidates(c *gin.Context) {
    projectIDStr := c.Query("projectId")
    if projectIDStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "缺少项目ID"})
        return
    }

    projectID, err := strconv.Atoi(projectIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的项目ID"})
        return
    }
    
    // 使用全局服务
    globalService := services.GetGlobalContractService()
    if globalService == nil {
        log.Println("错误: 全局服务未初始化")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "服务未初始化"})
        return
    }

    candidates, err := globalService.GetAllCandidates(uint(projectID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取候选人失败", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"candidates": candidates})
}