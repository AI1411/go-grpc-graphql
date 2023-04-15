package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
)

func TestUserStatusString(t *testing.T) {
	assert.Equal(t, "通常会員", entity.UserStatusActive.String())
	assert.Equal(t, "退会済", entity.UserStatusResigned.String())
	assert.Equal(t, "アカウント停止", entity.UserStatusBanded.String())
	assert.Equal(t, "プレミアム", entity.UserStatusPremium.String())
}

func TestNewUserStatus(t *testing.T) {
	userStatus := entity.NewUserStatus("通常会員")
	assert.Equal(t, entity.UserStatus("通常会員"), userStatus)

	userStatus = entity.NewUserStatus("退会済")
	assert.Equal(t, entity.UserStatus("退会済"), userStatus)

	userStatus = entity.NewUserStatus("アカウント停止")
	assert.Equal(t, entity.UserStatus("アカウント停止"), userStatus)

	userStatus = entity.NewUserStatus("プレミアム")
	assert.Equal(t, entity.UserStatus("プレミアム"), userStatus)
}

func TestUserStatusName(t *testing.T) {
	assert.Equal(t, entity.UserStatus("通常会員"), entity.UserStatusName["ACTIVE"])
	assert.Equal(t, entity.UserStatus("退会済"), entity.UserStatusName["RESIGNED"])
	assert.Equal(t, entity.UserStatus("アカウント停止"), entity.UserStatusName["BANDED"])
	assert.Equal(t, entity.UserStatus("プレミアム"), entity.UserStatusName["PREMIUM"])
}

func TestUserStatusValue(t *testing.T) {
	assert.Equal(t, entity.UserStatus("ACTIVE"), entity.UserStatusValue["通常会員"])
	assert.Equal(t, entity.UserStatus("RESIGNED"), entity.UserStatusValue["退会済"])
	assert.Equal(t, entity.UserStatus("BANDED"), entity.UserStatusValue["アカウント停止"])
	assert.Equal(t, entity.UserStatus("PREMIUM"), entity.UserStatusValue["プレミアム"])
}
