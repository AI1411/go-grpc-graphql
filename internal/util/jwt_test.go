package util

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	// テスト用のユーザーIDを作成
	userID := "testUser"

	// トークンを生成
	token, err := GenerateToken(userID)
	assert.NoError(t, err, "GenerateToken should not return an error")

	// 署名されたトークンを解析
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})
	assert.NoError(t, err, "Parsing the token should not return an error")

	// トークンが有効であることを確認
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok, "Token should have valid claims")
	assert.Equal(t, userID, claims["userId"], "Token should contain the correct user ID")

	// トークンの有効期限が正しいことを確認
	expirationTime := int64(claims["exp"].(float64))
	assert.Less(t, time.Now().Unix(), expirationTime, "Token should not be expired")
	assert.LessOrEqual(t, expirationTime, time.Now().Unix()+int64(24*time.Hour), "Token should have no more than 24 hours left")
}

func TestValidateJWT(t *testing.T) {
	userID := "test_user"

	// 生成されたトークンを検証します。
	token, err := GenerateToken(userID)
	assert.NoError(t, err)

	valid, err := ValidateJWT(token)
	assert.NoError(t, err)
	assert.True(t, valid)

	// 無効なトークンを検証します。
	invalidToken := "invalid_token"
	valid, err = ValidateJWT(invalidToken)
	assert.Error(t, err)
	assert.False(t, valid)
}
