package user

import "github.com/AI1411/go-grpc-graphql/grpc"

type GetUserHobbiesForm struct {
	UserID string `jaFieldName:"ユーザID" validate:"required,uuid4"`
}

func NewGetUserHobbiesForm(in *grpc.GetUserHobbiesRequest) *GetUserHobbiesForm {
	return &GetUserHobbiesForm{
		UserID: in.GetUserId(),
	}
}
