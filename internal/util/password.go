package util

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/AI1411/go-grpc-praphql/internal/domain/user/entity"
)

func SetPassword(user *entity.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(password)
	return nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
