package user

import "github.com/AI1411/go-grpc-graphql/grpc"

type CreateUserForm struct {
	Username     string          `jaFieldName:"ユーザ名" validate:"required,lte=100"`
	Email        string          `jaFieldName:"メールアドレス" validate:"required,email,lte=100"`
	Password     string          `jaFieldName:"パスワード" validate:"required,gte=8,lte=100"`
	Prefecture   grpc.Prefecture `jaFieldName:"都道府県" validate:"required,prefecture"`
	Introduction string          `jaFieldName:"自己紹介" validate:"required,lte=1000"`
	BloodType    grpc.BloodType  `jaFieldName:"血液型" validate:"required,bloodType"`
}

func NewCreateUserForm(in *grpc.CreateUserRequest) *CreateUserForm {
	return &CreateUserForm{
		Username:     in.GetUsername(),
		Email:        in.GetEmail(),
		Password:     in.GetPassword(),
		Prefecture:   in.GetPrefecture(),
		Introduction: in.GetIntroduction(),
		BloodType:    in.GetBloodType(),
	}
}
