package category

import "github.com/AI1411/go-grpc-graphql/grpc"

type ListCategoryForm struct {
	Name   string `jaFieldName:"name" validate:"required"`
	Order  string `jaFieldName:"order" validate:"omitempty"`
	Limit  int32  `jaFieldName:"limit" validate:"omitempty"`
	Offset int32  `jaFieldName:"offset" validate:"omitempty"`
}

func NewListCategoryForm(in *grpc.ListCategoryRequest) *ListCategoryForm {
	return &ListCategoryForm{
		Name:   in.GetName(),
		Order:  in.GetOrder(),
		Limit:  in.GetLimit(),
		Offset: in.GetOffset(),
	}
}
