package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加候选人
func AddCandidate(c *gin.Context) {
    var candidateData struct {
        ProjectID   uint   `json:"project_id" binding:"required"`
        Name        string `json:"name" binding:"required"`
        ImageUrl    string `json:"image_url" binding:"required"`
        Description string `json:"description" binding:"required"`
        // 移除 Address 字段
    }

    if err := c.ShouldBindJSON(&candidateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
        return
    }

    err := contractService.AddCandidate(
        candidateData.ProjectID,
        candidateData.Name,
        candidateData.ImageUrl,
        candidateData.Description,
        // 不要传递地址参数
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "添加候选人失败", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "候选人添加成功"})
}

// 修改候选人
func UpdateCandidate(c *gin.Context) {
    var candidateData struct {
        ProjectID   uint   `json:"project_id" binding:"required"`
        CandidateID uint   `json:"candidate_id" binding:"required"`
        Name        string `json:"name" binding:"required"`
        ImageUrl    string `json:"image_url" binding:"required"`
        Description string `json:"description" binding:"required"`
    }

    if err := c.ShouldBindJSON(&candidateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
        return
    }

    err := contractService.UpdateCandidate(
        candidateData.ProjectID,
        candidateData.CandidateID,
        candidateData.Name,
        candidateData.ImageUrl,
        candidateData.Description,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "修改候选人失败", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "候选人修改成功"})
}

// 删除候选人
func DeleteCandidate(c *gin.Context) {
    id := c.Param("id")
    projectID := c.Query("project_id")
    
    if id == "" || projectID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "缺少候选人ID或项目ID"})
        return
    }

    // 将ID转换为uint
    candidateID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "候选人ID格式错误"})
        return
    }
    
    projectIDUint, err := strconv.ParseUint(projectID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "项目ID格式错误"})
        return
    }

    err = contractService.DeleteCandidate(uint(projectIDUint), uint(candidateID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "删除候选人失败", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "候选人删除成功"})
}

// 开始投票
func StartVoting(c *gin.Context) {
    var data struct {
        ProjectID uint `json:"project_id" binding:"required"`
    }

    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
        return
    }

    err := contractService.StartVoting(data.ProjectID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "开始投票失败: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "投票已开始"})
}

// 结束投票
func EndVoting(c *gin.Context) {
    var data struct {
        ProjectID uint `json:"project_id" binding:"required"`
    }

    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
        return
    }

    err := contractService.EndVoting(data.ProjectID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "结束投票失败: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "投票已结束"})
}

// 创建投票项目 (新增功能)
func CreateVotingProject(c *gin.Context) {
    var projectData struct {
        Description string `json:"description" binding:"required"`
    }

    if err := c.ShouldBindJSON(&projectData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
        return
    }

    err := contractService.CreateVotingProject(projectData.Description)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "创建投票项目失败", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "投票项目创建成功"})
}

// 获取所有项目 (新增功能)
func GetProjects(c *gin.Context) {
    count, err := contractService.GetProjectCount()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取项目失败", "details": err.Error()})
        return
    }

    var projects []map[string]interface{}
    for i := uint(1); i <= count; i++ {
        project, err := contractService.GetProjectInfo(i)
        if err != nil {
            continue // 跳过获取失败的项目
        }
        projects = append(projects, project)
    }

    c.JSON(http.StatusOK, gin.H{"projects": projects})
}
