package hobby

import (
	"context"

	hobbyRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/hobby"
)

type DeleteHobbyUsecaseImpl interface {
	Exec(ctx context.Context, id string) error
}

type deleteHobbyUsecaseImpl struct {
	hobbyRepo hobbyRepo.Repository
}

func NewDeleteHobbyUsecaseImpl(hobbyRepo hobbyRepo.Repository) DeleteHobbyUsecaseImpl {
	return &deleteHobbyUsecaseImpl{
		hobbyRepo: hobbyRepo,
	}
}

func (u *deleteHobbyUsecaseImpl) Exec(ctx context.Context, id string) error {
	return u.hobbyRepo.DeleteHobby(ctx, id)
}
