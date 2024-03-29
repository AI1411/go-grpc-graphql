package util

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
)

const cryptCost = 10

func SetPassword(user *entity.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), cryptCost)
	if err != nil {
		return err
	}
	user.Password = string(password)
	return nil
}
