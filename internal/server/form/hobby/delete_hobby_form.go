package hobby

import "github.com/AI1411/go-grpc-graphql/grpc"

type DeleteHobbyForm struct {
	ID string `jaFieldName:"ID" validate:"required"`
}

func NewDeleteHobbyForm(in *grpc.DeleteHobbyRequest) *DeleteHobbyForm {
	return &DeleteHobbyForm{
		ID: in.GetId(),
	}
}
