package user

import "github.com/AI1411/go-grpc-praphql/grpc"

type LoginForm struct {
	Email    string `jaFieldName:"メールアドレス" validate:"required,email"`
	Password string `jaFieldName:"パスワード" validate:"required"`
}

func NewLoginForm(in *grpc.LoginRequest) *LoginForm {
	return &LoginForm{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}
}
