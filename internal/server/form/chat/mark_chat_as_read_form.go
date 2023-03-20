package chat

import "github.com/AI1411/go-grpc-graphql/grpc"

type MarkChatAsReadForm struct {
	ID string `jaFieldName:"ID" validate:"required,uuid4"`
}

func NewMarkChatAsReadForm(in *grpc.MarkChatAsReadRequest) *MarkChatAsReadForm {
	return &MarkChatAsReadForm{
		ID: in.GetId(),
	}
}
