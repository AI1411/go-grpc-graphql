package user

import "github.com/AI1411/go-grpc-graphql/grpc"

type GetUserForm struct {
	ID string `jaFieldName:"ユーザID" validate:"required,uuid4"`
}

func NewGetUserForm(in *grpc.GetUserRequest) *GetUserForm {
	return &GetUserForm{
		ID: in.GetId(),
	}
}
