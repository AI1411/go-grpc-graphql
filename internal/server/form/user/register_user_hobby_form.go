package user

import "github.com/AI1411/go-grpc-graphql/grpc"

type RegisterUserHobbyForm struct {
	UserID  string `jaFieldName:"ユーザID" validate:"required"`
	HobbyID string `jaFieldName:"趣味ID" validate:"required"`
}

func NewRegisterUserHobbyForm(in *grpc.RegisterUserHobbyRequest) *RegisterUserHobbyForm {
	return &RegisterUserHobbyForm{
		UserID:  in.GetUserId(),
		HobbyID: in.GetHobbyId(),
	}
}
