package controllers

import (
	"net/http"
	"product-store/app/models"
	"product-store/app/resources"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// GetUserByToken retrieves the user information based on the JWT token
func (u *UserController) GetUserByToken(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var user models.User
	if err := u.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	userResource := resources.UserResource(user)
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "User fetched successfully",
		"user":    userResource,
	})
}
