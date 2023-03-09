package util

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/AI1411/go-grpc-praphql/internal/domain/user/entity"
)

func SetPassword(user *entity.User) error {
	password, err := bcrypt.GenerateFromPassword(user.Password, 10)
	if err != nil {
		return err
	}
	user.Password = password
	return nil
}

func ComparePassword(user *entity.User, password []byte) error {
	return bcrypt.CompareHashAndPassword(user.Password, password)
}
