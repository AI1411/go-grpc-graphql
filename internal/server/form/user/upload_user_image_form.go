package user

import "github.com/AI1411/go-grpc-graphql/grpc"

type UploadUserImageForm struct {
	UserId string `jaFieldName:"ユーザID" validate:"required"`
	Image  string `jaFieldName:"画像" validate:"required"`
}

func NewUploadUserImageForm(in *grpc.UploadUserImageRequest) *UploadUserImageForm {
	return &UploadUserImageForm{
		UserId: in.GetUserId(),
		Image:  in.GetImage(),
	}
}
