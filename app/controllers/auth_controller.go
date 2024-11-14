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
	DB      *gorm.DB
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

// Login handles user login and token generation
func (a *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate user credentials (this is just an example, implement your logic)
	var foundUser models.User
	if err := a.DB.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Here you should verify the password, using a function like HashPassword()
	if err := utils.VerifyPassword(foundUser.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(foundUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Return the token to the client
	c.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"message":     "Successfully generated token",
		"token":       token,
		"accessToken": token,
	})
}
func (a *AuthController) Logout(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Delete the session from the database
	a.DB.Where("refresh_token = ?", request.RefreshToken).Delete(&models.Session{})
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
