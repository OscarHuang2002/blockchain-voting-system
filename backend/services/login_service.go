package services

import (
	"backend/models"
	"backend/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService struct {
	DB *gorm.DB
}

func NewLoginService(db *gorm.DB) *LoginService {
	return &LoginService{DB: db}
}

// 使用钱包地址登录
func (service *LoginService) LoginWithAddress(address string, password string) (string, error) {
	var user models.User
	if err := service.DB.Where("address = ?", address).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("用户未注册，请先注册")
		}
		return "", fmt.Errorf("failed to query user: %v", err)
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("密码错误")
	}

	// 生成 JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

func (service *LoginService) IsAdmin(address string) (bool, error) {
	var user models.User
	if err := service.DB.Where("address = ?", address).First(&user).Error; err != nil {
		return false, err
	}

	// 检查是否为管理员地址
	return address == "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", nil
}
