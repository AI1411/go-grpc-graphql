package chat

import "github.com/AI1411/go-grpc-graphql/grpc"

type CreateChatForm struct {
	roomID     string `jaFieldName:"ルームID" validate:"required,uuid4"`
	fromUserID string `jaFieldName:"送信者ユーザID" validate:"required,uuid4"`
	toUserID   string `jaFieldName:"受信者ユーザID" validate:"required,uuid4"`
	body       string `jaFieldName:"本文" validate:"required. max=255"`
}

func NewCreateChatForm(in *grpc.CreateChatRequest) *CreateChatForm {
	return &CreateChatForm{
		roomID:     in.GetRoomId(),
		fromUserID: in.GetFromUserId(),
		toUserID:   in.GetToUserId(),
		body:       in.GetBody(),
	}
}
