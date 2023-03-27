package user

import "github.com/AI1411/go-grpc-graphql/grpc"

type DeleteUserHobbyForm struct {
	UserID  string `jaFieldName:"ユーザID" validate:"required"`
	HobbyID string `jaFieldName:"趣味ID" validate:"required"`
}

func NewDeleteUserHobbyForm(in *grpc.DeleteUserHobbyRequest) *DeleteUserHobbyForm {
	return &DeleteUserHobbyForm{
		UserID:  in.GetUserId(),
		HobbyID: in.GetHobbyId(),
	}
}
