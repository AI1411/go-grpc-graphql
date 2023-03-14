package user

import "github.com/AI1411/go-grpc-graphql/grpc"

type UpdateUserPasswordForm struct {
	ID                   string `jaFieldName:"ユーザID" validate:"required,uuid4"`
	Password             string `jaFieldName:"新しいパスワード" validate:"required,lte=100,nefield=ExPassword"`
	PasswordConfirmation string `jaFieldName:"新しいパスワード(確認)" validate:"required,lte=100,eqfield=Password"`
}

func NewUpdateUserPasswordForm(in *grpc.UpdateUserPasswordRequest) *UpdateUserPasswordForm {
	return &UpdateUserPasswordForm{
		ID:                   in.GetId(),
		Password:             in.GetPassword(),
		PasswordConfirmation: in.GetPasswordConfirmation(),
	}
}
