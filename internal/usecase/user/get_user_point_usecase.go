package user

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
)

type GetUserPointUsecase interface {
	Exec(ctx context.Context, userID string) (*grpc.GetUserPointResponse, error)
}

type getUserPointUsecaseImpl struct {
	userRepository      userRepo.Repository
	userPointRepository userRepo.PointRepository
}

func NewGetUserPointUsecaseImpl(userRepository userRepo.Repository, userPointRepository userRepo.PointRepository) GetUserPointUsecase {
	return &getUserPointUsecaseImpl{
		userRepository:      userRepository,
		userPointRepository: userPointRepository,
	}
}

func (usecase *getUserPointUsecaseImpl) Exec(ctx context.Context, userID string) (*grpc.GetUserPointResponse, error) {
	point, err := usecase.userPointRepository.GetPoint(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &grpc.GetUserPointResponse{Point: int32(point)}, nil
}
