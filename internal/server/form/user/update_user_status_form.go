package user

import "github.com/AI1411/go-grpc-praphql/grpc"

type UpdateUserStatusForm struct {
	ID     string      `jaFieldName:"ユーザID" validate:"required,uuid4"`
	Status grpc.Status `jaFieldName:"ユーザステータス" validate:"required"`
}

func NewUpdateUserStatusForm(in *grpc.UpdateUserStatusRequest) *UpdateUserStatusForm {
	return &UpdateUserStatusForm{
		ID:     in.GetId(),
		Status: in.GetStatus(),
	}
}
