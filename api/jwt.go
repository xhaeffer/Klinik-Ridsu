package jwt

import (
//	"encoding/base64"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

var (
	// secretKey adalah kunci rahasia untuk JWT
	secretKey = []byte("klinikridsu")
	
	// ErrMissingToken adalah error untuk token JWT yang hilang
	ErrMissingToken = errors.New("missing JWT token")
	
	// ErrTokenInvalid adalah error untuk token JWT yang tidak valid
	ErrTokenInvalid = errors.New("invalid JWT token")
)

// GenerateToken menghasilkan token JWT baru
func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Waktu kedaluwarsa token (1 hari)
	// Tambahkan klaim lebih banyak sesuai kebutuhan

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken memverifikasi token JWT
func VerifyToken(c *gin.Context) error {
	tokenString, err := c.Cookie("token")
	if err != nil {
		return ErrMissingToken
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return ErrTokenInvalid
	}

	return nil
}
