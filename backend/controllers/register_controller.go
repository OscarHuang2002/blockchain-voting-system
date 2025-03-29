package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Global variables for contract service and database instance
var contractService *services.ContractService
var db *gorm.DB

// Initialize dependencies for contract service and database
func InitDependencies(service *services.ContractService, database *gorm.DB) {
	contractService = service
	db = database
}

// Struct to handle incoming registration requests
type RegisterRequest struct {
	Username string `json:"username" binding:"required"` // Username, required
	Email    string `json:"email" binding:"required"`    // Email, required
	Password string `json:"password" binding:"required"` // Password, required
	Address  string `json:"address" binding:"required"`  // Blockchain address, required
}

// Controller function to register a voter
func RegisterVoter(c *gin.Context) {
	// Bind JSON data to the request struct
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash the password for security
	hashedPassword := utils.HashPassword(req.Password)

	// Save user information to the database
	user := models.User{
		Name:     req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Address:  req.Address,
	}

	// Check for duplicate users in the database
	if result := db.Create(&user); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Register the voter on the blockchain
	err := contractService.RegisterVoter(req.Address)
	if err != nil {
		// If blockchain registration fails, delete the saved user record
		db.Delete(&user)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Blockchain registration failed: " + err.Error()})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
