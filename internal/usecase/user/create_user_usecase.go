package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	commonEntity "github.com/AI1411/go-grpc-graphql/internal/domain/common/entity"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
)

type CreateUserUsecase interface {
	Exec(ctx context.Context, in *grpc.CreateUserRequest) error
}

type createUserUsecaseImpl struct {
	userRepository userRepo.UserRepository
}

func NewCreateUserUsecaseImpl(userRepository userRepo.UserRepository) CreateUserUsecase {
	return &createUserUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *createUserUsecaseImpl) Exec(ctx context.Context, in *grpc.CreateUserRequest) error {
	return u.userRepository.CreateUser(ctx, &entity.User{
		Username:     in.GetUsername(),
		Email:        in.GetEmail(),
		Status:       entity.UserStatusActive,
		Prefecture:   commonEntity.PrefectureName[in.GetPrefecture().String()],
		Introduction: in.GetIntroduction(),
		BloodType:    commonEntity.BloodTypeName[in.GetBloodType().String()],
	})
}
