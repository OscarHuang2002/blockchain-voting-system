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

func (service *LoginService) LoginUser(email, password string) (string, error) {
	var user models.User
	if err := service.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", fmt.Errorf("invalid email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password")
	}

	// 生成 JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}
