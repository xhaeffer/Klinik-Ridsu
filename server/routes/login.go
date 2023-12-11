package routes

import (
	// "crypto/sha256"
	// "encoding/hex"
	"net/http"
	"strconv"
	"KlinikRidsu/server/hash"
	"KlinikRidsu/server/databases"
	"KlinikRidsu/server/session"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Identifier int `form:"identifier" binding:"required"`
	Password   string `form:"password" binding:"required"`
}

func Login(r *gin.Engine, db *gorm.DB) {
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		var loginRequest LoginRequest

		if err := c.ShouldBind(&loginRequest); err != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Invalid request payload"})
			return
		}

		var user databases.User
		if err := db.Where("no_rs = ? OR nik = ?", checkIdentifier(strconv.Itoa(loginRequest.Identifier)), checkIdentifier(strconv.Itoa(loginRequest.Identifier))).First(&user).Error; err != nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "No RS / NIK / Password anda salah"})
			return
		}

		if !verifyPassword(loginRequest.Password, user.Password) {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "No RS / NIK / Password anda salah"})
			return
		}

		err := session.SetSession(c.Writer, c.Request, "user", user)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Gagal menyimpan sesi"})
			return
		}

		token, err := session.GenerateToken()
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Failed to generate token"})
			return
		}

		tokenString, ok := token["token"].(string)
		if !ok {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Failed to get token string"})
			return
		}

		c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": "Login successful",
			"token":   token,
		})
	})
}
func checkIdentifier(identifier string) string {
    if len(identifier) == 16 {
        hashed, err := hash.HashNIK(identifier)
        if err != nil {
            return identifier
        }
        return hashed
    }
    return identifier
}

func verifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
