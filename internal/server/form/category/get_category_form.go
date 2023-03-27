package category

import "github.com/AI1411/go-grpc-graphql/grpc"

type GetCategoryForm struct {
	ID string `jaFieldName:"ID" validate:"required"`
}

func NewGetCategoryForm(in *grpc.GetCategoryRequest) *GetCategoryForm {
	return &GetCategoryForm{
		ID: in.GetId(),
	}
}
