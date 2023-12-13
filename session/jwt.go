package session

import (
//	"encoding/base64"
	"net/http"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = []byte("klinikridsu")
	ErrMissingToken = errors.New("missing JWT token")
	ErrTokenInvalid = errors.New("invalid JWT token")
)

func GenerateToken() (map[string]interface{}, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)

    expirationTime := time.Now().Add(time.Hour * 24)
    claims["exp"] = expirationTime.Unix()

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return nil, err
    }

    expiresIn := time.Until(expirationTime)

    additionalInfo := map[string]interface{}{
        "token":       tokenString,
        "expires_in":  expiresIn.String(),
        "user_id":     "admin",
    }
    return additionalInfo, nil
}

func ExtractToken(c *gin.Context) error {
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

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := ExtractToken(c); err != nil {
			if err == ErrMissingToken {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan, Silahkan Login terlebih dahulu!"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid, Silahkan login ulang!"})
			}
			c.Abort()
			return
		}
	}
}