package controllers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var loginService *services.LoginService

func InitLoginService(database *gorm.DB) {
	loginService = services.NewLoginService(database)
}

// 使用钱包地址和密码登录
func LoginWithAddress(c *gin.Context) {
	var loginData struct {
		Address string `json:"address" binding:"required"`
		Passwd  string `json:"passwd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := loginService.LoginWithAddress(loginData.Address, loginData.Passwd)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": token})
}
