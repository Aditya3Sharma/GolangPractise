package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("aditya")

func CreateToken(username string) (string, error) {
	fmt.Println(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	// fmt.Println(token)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	// fmt.Println(tokenString)
	// fmt.Println(VerifyToken(tokenString))
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	// fmt.Println("Token after parsing:", token)
	// fmt.Println("Jingalala huhu")
	fmt.Println(secretKey)
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
