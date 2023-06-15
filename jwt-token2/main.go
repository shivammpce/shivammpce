package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	KEY_ERROR     = "Error in generating secret key!"
	TOKEN_ERROR   = "Error in generating JWT token!"
	SUCCESS_KEY   = "Successfully Generating key!"
	SUCCESS_TOKEN = "Successfully Generating token!"
	ERROR         = "Something went wrong"
)

func main() {
	// generate a random secret key
	SecKey, err := generateRandomKey(32)
	if err != nil {
		fmt.Println(KEY_ERROR, err)
		return
	}

	// if len(SecKey) >= 32 {
	// 	fmt.Println(SUCCESS_KEY, SecKey)
	// } else if len(SecKey) <= 32 {
	// 	fmt.Println(ERROR, SecKey)
	// }

	fmt.Println(SUCCESS_KEY, SecKey)

	// generate the JWT token
	TokStr, err := generateToken(SecKey)
	if err != nil {
		fmt.Println(TOKEN_ERROR, err)
		return
	}

	fmt.Println(SUCCESS_TOKEN, TokStr)

}

func generateRandomKey(keyLength int) (string, error) {
	key := make([]byte, keyLength)

	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func generateToken(secretKey string) (string, error) {

	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// set the claims (payload) of the token
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = 123
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// generate the token string
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return string(tokenString), nil

}
