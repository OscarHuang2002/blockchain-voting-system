package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 上传图片
func UploadImage(c *gin.Context) {
    // 获取上传的文件
    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "获取上传文件失败", "details": err.Error()})
        return
    }

    // 创建上传目录
    uploadDir := "./uploads/images"
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        if err := os.MkdirAll(uploadDir, 0755); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "创建上传目录失败", "details": err.Error()})
            return
        }
    }

    // 生成唯一文件名
    fileExt := filepath.Ext(file.Filename)
    fileName := fmt.Sprintf("%s_%s%s", time.Now().Format("20060102150405"), uuid.New().String(), fileExt)
    filePath := filepath.Join(uploadDir, fileName)

    // 保存文件
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败", "details": err.Error()})
        return
    }

    // 返回可访问的图片URL
    imageURL := fmt.Sprintf("http://localhost:8080/images/%s", fileName)
    c.JSON(http.StatusOK, gin.H{"url": imageURL})
}