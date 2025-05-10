package controllers

import (
	"net/http"
	"traceability/models"
	"traceability/services"

	"github.com/gin-gonic/gin"
)

var validUser = models.User{
	Username: "admin",
	Password: "123456", // Đổi sang hash thật nếu cần
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if user.Username != validUser.Username || user.Password != validUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong credentials"})
		return
	}

	token, err := services.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
