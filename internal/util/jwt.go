package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("my_secret_key")

func GenerateToken(userID string) (string, error) {
	// JWTに含めるクレーム（Claim）の設定
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     userID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // トークンの有効期限は24時間
	}

	// JWTの設定
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンの署名
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if !claims["authorized"].(bool) {
			return false, nil
		}
		return true, nil
	}

	return false, err
}

func GetUserIDFromJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["userId"].(string), nil
	}

	return "", err
}
