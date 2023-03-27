package category

import "github.com/AI1411/go-grpc-graphql/grpc"

type CreateCategoryForm struct {
	Name        string `jaFieldName:"カテゴリ名" validate:"required"`
	Description string `jaFieldName:"カテゴリ説明" validate:"required"`
}

func NewCreateCategoryForm(in *grpc.CreateCategoryRequest) *CreateCategoryForm {
	return &CreateCategoryForm{
		Name:        in.GetName(),
		Description: in.GetDescription(),
	}
}
