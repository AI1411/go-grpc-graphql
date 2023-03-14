package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	commonEntity "github.com/AI1411/go-grpc-graphql/internal/domain/common/entity"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	repository "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type UpdateUserProfileUsecase interface {
	Exec(ctx context.Context, in *grpc.UpdateUserProfileRequest) error
}

type updateUserProfileUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUpdateUserProfileUsecaseImpl(userRepository repository.UserRepository) UpdateUserProfileUsecase {
	return &updateUserProfileUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u updateUserProfileUsecaseImpl) Exec(ctx context.Context, in *grpc.UpdateUserProfileRequest) error {
	return u.userRepository.UpdateUserProfile(ctx, &entity.User{
		ID:           util.StringToNullUUID(in.GetId()),
		Username:     in.GetUsername(),
		Prefecture:   commonEntity.PrefectureName[in.GetPrefecture().String()],
		Introduction: in.GetIntroduction(),
		BloodType:    commonEntity.BloodTypeName[in.GetBloodType().String()],
	})
}
