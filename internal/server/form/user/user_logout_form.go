package user

import "github.com/AI1411/go-grpc-graphql/grpc"

type LogoutForm struct {
	Token  string `jaFieldName:"トークン" validate:"required"`
	UserID string `jaFieldName:"ユーザID" validate:"required,uuid4"`
}

func NewLogoutForm(in *grpc.LogoutRequest) *LogoutForm {
	return &LogoutForm{
		Token:  in.GetToken(),
		UserID: in.GetUserId(),
	}
}
