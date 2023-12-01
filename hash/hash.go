package hash

import (
    "crypto/sha256"
    "encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func HashNIK(data string) (string, error) {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	hashedData := hex.EncodeToString(hasher.Sum(nil))
	return hashedData, nil
}

func HashPassword(text string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}