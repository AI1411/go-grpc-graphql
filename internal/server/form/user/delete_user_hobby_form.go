package user

import "github.com/AI1411/go-grpc-graphql/grpc"

type DeleteUserHobbyForm struct {
	ID string `jaFieldName:"ID" validate:"required"`
}

func NewDeleteUserHobbyForm(in *grpc.DeleteUserHobbyRequest) *DeleteUserHobbyForm {
	return &DeleteUserHobbyForm{
		ID: in.GetId(),
	}
}
