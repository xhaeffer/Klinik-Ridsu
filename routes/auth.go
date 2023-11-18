package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"projek1/api"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Auth(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", LoginHandler)
		authGroup.POST("/logout", LogoutHandler)
	}
}

func LoginHandler(c *gin.Context) {
	var loginRequest LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if loginRequest.Username != "ridsa" || loginRequest.Password != "ridsu" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := jwt.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
