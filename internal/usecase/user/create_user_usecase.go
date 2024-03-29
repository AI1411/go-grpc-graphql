package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/AI1411/go-grpc-graphql/grpc"
	commonEntity "github.com/AI1411/go-grpc-graphql/internal/domain/common/entity"
	"github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
)

type CreateUserUsecase interface {
	Exec(ctx context.Context, in *grpc.CreateUserRequest) (*grpc.CreateUserResponse, error)
}

type createUserUsecaseImpl struct {
	userRepository userRepo.Repository
}

func NewCreateUserUsecaseImpl(userRepository userRepo.Repository) CreateUserUsecase {
	return &createUserUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *createUserUsecaseImpl) Exec(ctx context.Context, in *grpc.CreateUserRequest) (*grpc.CreateUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	res, err := u.userRepository.CreateUser(ctx, &entity.User{
		Username:     in.GetUsername(),
		Email:        in.GetEmail(),
		Password:     string(hashedPassword),
		Status:       entity.UserStatusActive,
		Prefecture:   commonEntity.PrefectureName[in.GetPrefecture().String()],
		Introduction: in.GetIntroduction(),
		BloodType:    commonEntity.BloodTypeName[in.GetBloodType().String()],
	})
	if err != nil {
		return nil, err
	}

	return &grpc.CreateUserResponse{Id: res}, nil
}
