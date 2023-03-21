package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
)

func TestSetPassword(t *testing.T) {
	// テスト用のユーザーを作成
	user := &entity.User{
		Password: "test_password",
	}

	// パスワードを設定
	err := SetPassword(user)
	assert.NoError(t, err, "SetPassword should not return an error")

	// ハッシュ化されたパスワードが正しいことを確認
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("test_password"))
	assert.NoError(t, err, "The hashed password should match the original password")
}
