// app/controllers/auth_controller.go
package controllers

import (
	"net/http"
	"product-store/app/models"
	"product-store/app/requests"
	"product-store/app/services"
	"product-store/app/utils"
	"product-store/app/validators"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	Service *services.AuthService
	DB      *gorm.DB // Database instance
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{Service: service}
}

// Register handles user registration
func (a *AuthController) Register(c *gin.Context) {
	var request requests.RegisterRequest

	// Bind JSON body to RegisterRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate request
	errors := validators.ValidateStruct(&request)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	// Hash password
	hashedPassword, hashErr := utils.HashPassword(request.Password)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user model
	user := models.User{Name: request.Name, Email: request.Email, Password: hashedPassword}

	// Attempt to create user in the database
	if err := a.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// Login handles user login
func (a *AuthController) Login(c *gin.Context) {
	var request requests.LoginRequest

	// Bind JSON body to LoginRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Find user by email
	var user models.User
	if err := a.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if !utils.CheckPasswordHash(request.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token (this part requires your JWT utility functions)
	token, err := utils.GenerateJWT(user.ID) // Assuming GenerateJWT takes a user ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
