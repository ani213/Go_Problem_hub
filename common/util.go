package common

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
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

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email address"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	default:
		return "Invalid value"
	}
}

func VerifyToken(tokenString string, secret string) (jwt.MapClaims, error) {
	// Parse the token with a secret and check signing method
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure it's an HMAC token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate and extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token or claims")
}
