package hobby

import (
	"github.com/AI1411/go-grpc-graphql/grpc"
)

type CreateHobbyForm struct {
	Name        string `jaFieldName:"趣味名" validate:"required"`
	Description string `jaFieldName:"趣味説明""`
	CategoryID  string `jaFieldName:"カテゴリID" validate:"required"`
}

func NewCreateHobbyForm(in *grpc.CreateHobbyRequest) *CreateHobbyForm {
	return &CreateHobbyForm{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		CategoryID:  in.GetCategoryId(),
	}
}
