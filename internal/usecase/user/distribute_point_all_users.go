package user

import (
	"context"

	userRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/user"
)

type DistributePointAllUsers interface {
	Exec(context.Context, string) error
}

type distributePointAllUsersImpl struct {
	userRepository userRepo.Repository
	pointRepo      userRepo.PointRepository
}

func NewDistributePointAllUsersImpl(userRepository userRepo.Repository, pointRepo userRepo.PointRepository) DistributePointAllUsers {
	return &distributePointAllUsersImpl{
		userRepository: userRepository,
		pointRepo:      pointRepo,
	}
}

func (u *distributePointAllUsersImpl) Exec(ctx context.Context, reason string) error {
	err := u.pointRepo.DistributePointAllUsers(ctx, "")
	if err != nil {
		return err
	}
	return nil
}
