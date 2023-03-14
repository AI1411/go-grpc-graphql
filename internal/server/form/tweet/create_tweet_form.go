package tweet

import "github.com/AI1411/go-grpc-praphql/grpc"

type CreateTweetForm struct {
	Body   string `json:"body" validate:"required,max=140"`
	UserID string `json:"user_id" validate:"required,uuid4"`
}

func NewCreateTweetForm(in *grpc.CreateTweetRequest) *CreateTweetForm {
	return &CreateTweetForm{
		Body:   in.GetBody(),
		UserID: in.GetUserId(),
	}
}
