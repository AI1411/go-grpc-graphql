package user

import "github.com/AI1411/go-grpc-graphql/grpc"

type UpdateUserProfileForm struct {
	ID           string          `jaFieldName:"ユーザID" validate:"required,uuid4"`
	Username     string          `jaFieldName:"ユーザ名" validate:"required,lte=100"`
	Prefecture   grpc.Prefecture `jaFieldName:"都道府県" validate:"required,prefecture"`
	Introduction string          `jaFieldName:"自己紹介" validate:"required,lte=1000"`
	BloodType    grpc.BloodType  `jaFieldName:"血液型" validate:"required,bloodType"`
}

func NewUpdateUserProfileForm(in *grpc.UpdateUserProfileRequest) *UpdateUserProfileForm {
	return &UpdateUserProfileForm{
		ID:           in.GetId(),
		Username:     in.GetUsername(),
		Prefecture:   in.GetPrefecture(),
		Introduction: in.GetIntroduction(),
		BloodType:    in.GetBloodType(),
	}
}
