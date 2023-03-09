package user

import (
	"context"

	"github.com/AI1411/go-grpc-praphql/grpc"
	commonEntity "github.com/AI1411/go-grpc-praphql/internal/domain/common/entity"
	"github.com/AI1411/go-grpc-praphql/internal/domain/user/entity"
	repository "github.com/AI1411/go-grpc-praphql/internal/infra/repository/user"
	"github.com/AI1411/go-grpc-praphql/internal/util"
)

type Exec interface {
	Exec(ctx context.Context, in *grpc.UpdateUserProfileRequest) error
}

type updateUserProfileUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUpdateUserProfileUsecaseImpl(userRepository repository.UserRepository) Exec {
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
