package chat

import "github.com/AI1411/go-grpc-graphql/grpc"

type ListChatForm struct {
	RoomID     string `jaFieldName:"ルームID" validate:"required,uuid"`
	FromUserID string `jaFieldName:"送信者ユーザID" validate:"required,uuid"`
}

func NewListChatForm(in *grpc.ListChatRequest) *ListChatForm {
	return &ListChatForm{
		RoomID:     in.GetRoomId(),
		FromUserID: in.GetUserId(),
	}
}
