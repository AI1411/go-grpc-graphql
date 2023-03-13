package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID string) (string, error) {
	// JWTに含めるクレーム（Claim）の設定
	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // トークンの有効期限は24時間
	}

	// JWTの設定
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("my_secret_key") // 秘密鍵

	// トークンの署名
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
