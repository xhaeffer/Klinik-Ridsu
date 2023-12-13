package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"KlinikRidsu/databases"
	"KlinikRidsu/hash"
)

type RegisterRequest struct {
	NoRS     int `form:"no_rs"`
	NIK       string `form:"nik" binding:"required"`
	Password  string `form:"password" binding:"required"`
}

func Register(r *gin.Engine, db *gorm.DB) {
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	r.POST("/register", func(c *gin.Context) {
		var registerRequest RegisterRequest

		if err := c.ShouldBind(&registerRequest); err != nil {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Invalid request payload"})
			return
		}

		hashedPassword, err := hash.HashPassword(registerRequest.Password)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Failed to hash password"})
			return
		}

		hashedNIK, err := hash.HashNIK(registerRequest.NIK)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Failed to hash NIK"})
			return
		}

		newUser := databases.User{
			NIK:      hashedNIK,
			Password: hashedPassword,
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Failed to create user"})
			return
		}

		c.HTML(http.StatusOK, "register.html", gin.H{
			"message": "Registration successful",
		})
	})
}