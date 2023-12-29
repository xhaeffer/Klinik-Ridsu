package login

import (
	"net/http"
	"strconv"
	"KlinikRidsu/databases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Identifier int `form:"identifier" binding:"required"`
	Password   string `form:"password" binding:"required"`
}

func Login(r *gin.Engine, db *gorm.DB) {
	r.POST("/login", func(c *gin.Context) {
		if isUserLoggedIn(c) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid!"})
			return
		}
		
		var loginRequest LoginRequest
		if err := c.ShouldBind(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid!"})
			return
		}

		var user databases.User
		if err := db.Where("no_rs = ? OR nik = ?", checkIdentifier(strconv.Itoa(loginRequest.Identifier)), checkIdentifier(strconv.Itoa(loginRequest.Identifier))).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Akun tidak terdaftar!"})
			return
		}

		if !verifyPassword(loginRequest.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Password anda salah!"})
			return
		}

		loginHandler(c, db, user)
	})

	
}