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

// GenerateTokenWithInfo generates a JWT token with expiration information
func GenerateToken() (map[string]interface{}, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)

    // Set expiration time for the token (1 day)
    expirationTime := time.Now().Add(time.Hour * 24)
    claims["exp"] = expirationTime.Unix()

    // Add more claims as needed

    // Sign and get the signed token string
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return nil, err
    }

    // Calculate the remaining time until expiration using time.Until
    expiresIn := time.Until(expirationTime)

    // Information to be returned along with the token
    additionalInfo := map[string]interface{}{
        "token":       tokenString,
        "expires_in":  expiresIn.String(), // Format the duration as a string
        "user_id":     "admin",
    }

    return additionalInfo, nil
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
