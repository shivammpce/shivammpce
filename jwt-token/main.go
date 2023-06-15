package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// Generate a random secret key
	secretKey := generateRandomSecretKey(32)
	fmt.Println("Secret Key:", secretKey)

	// Create a JWT token
	tokenString, err := createJWTToken(secretKey)
	if err != nil {
		log.Fatal("Failed to create JWT token:", err)
	}

	fmt.Println("JWT Token:", tokenString)

	// Verify the JWT token
	// claims, err := verifyJWTToken(tokenString, secretKey)
	// if err != nil {
	// 	log.Fatal("Failed to verify JWT token:", err)
	// }

	// fmt.Println("Claims:", claims)
}

// Generates a random secret key of the specified length
func generateRandomSecretKey(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

// Creates a JWT token using the provided secret key
func createJWTToken(secretKey string) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"sub": "1234567890",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verifies and parses the provided JWT token using the secret key
// func verifyJWTToken(tokenString, secretKey string) (*jwt.MapClaims, error) {
// 	// Parse the token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(secretKey), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Verify the token's signature algorithm
// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 	}

// 	// Extract the claims
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		return nil, fmt.Errorf("invalid JWT token")
// 	}

// 	return &claims, nil
// }
