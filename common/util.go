package common

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/pbkdf2"
)

func EncryptPassword(password, salt string) string {
	iterations := 10000
	keyLength := 64 // 64 bytes (512 bits)

	derivedKey := pbkdf2.Key([]byte(password), []byte(salt), iterations, keyLength, sha512.New)
	return hex.EncodeToString(derivedKey) // Convert to hex string
}

func CheckPassword(password string, hashpassword string, salt string) bool {
	return EncryptPassword(password, salt) == hashpassword

}

// GenerateAccessToken creates a JWT access token
func GenerateAccessToken(data map[string]interface{}) (string, error) {
	secretKey := os.Getenv("TOKEN_KEY") // Get the secret key from environment variables
	return generateToken(data, secretKey, 7*time.Hour)
}

// GenerateRefreshToken creates a JWT refresh token
func GenerateRefreshToken(data map[string]interface{}) (string, error) {
	secretKey := os.Getenv("TOKEN_KEY") + "REFRESH"
	return generateToken(data, secretKey, 7*time.Hour)
}

// Common function to generate JWT tokens
func generateToken(data map[string]interface{}, secretKey string, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(expiration).Unix(), // Set expiration time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func GenerateSalt() (string, error) {
	bytes := make([]byte, 32)  // Create a 32-byte slice
	_, err := rand.Read(bytes) // Fill with cryptographic random bytes
	if err != nil {
		return "", err // Handle error if random generation fails
	}
	return hex.EncodeToString(bytes), nil // Convert to hex string
}
