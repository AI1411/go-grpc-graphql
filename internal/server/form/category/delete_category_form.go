package category

import "github.com/AI1411/go-grpc-graphql/grpc"

type DeleteCategoryForm struct {
	ID string `jajaFieldName:"ID" validate:"required"`
}

func NewDeleteCategoryForm(in *grpc.DeleteCategoryRequest) *DeleteCategoryForm {
	return &DeleteCategoryForm{
		ID: in.GetId(),
	}
}
