package register

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
	r.POST("/register", func(c *gin.Context) {
		var registerRequest RegisterRequest

		if err := c.ShouldBind(&registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid!"})
			return
		}

		hashedNIK, err := hash.HashNIK(registerRequest.NIK)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash NIK!"})
			return
		}

		hashedPassword, err := hash.HashPassword(registerRequest.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password!"})
			return
		}

		newUser := databases.User{
			NIK:      hashedNIK,
			Password: hashedPassword,
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat akun!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{ "message": "Registrasi Berhasil!",})
	})
}