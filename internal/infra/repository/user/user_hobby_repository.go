package user

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-graphql/internal/domain/hobby/entity"
	userEntity "github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type UserHobbyRepository interface {
	GetUserHobbies(context.Context, string) ([]*entity.Hobby, error)
	RegisterUserHobby(context.Context, *entity.Hobby) (string, error)
	DeleteUserHobby(context.Context, string) error
}

type userHobbyRepository struct {
	dbClient *db.Client
}

func NewUserHobbyRepository(dbClient *db.Client) UserHobbyRepository {
	return &userHobbyRepository{
		dbClient: dbClient,
	}
}

func (u userHobbyRepository) GetUserHobbies(ctx context.Context, userID string) ([]*entity.Hobby, error) {
	var user userEntity.User
	if err := u.dbClient.Conn(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, err
	}

	var userHobbies []*userEntity.UserHobby
	if err := u.dbClient.Conn(ctx).Where("user_id", userID).Find(&userHobbies).Error; err != nil {
		return nil, err
	}

	if len(userHobbies) == 0 {
		return nil, nil
	}

	hobbyIDs := make([]string, len(userHobbies))
	for i, hobby := range userHobbies {
		hobbyIDs[i] = util.NullUUIDToString(hobby.HobbyID)
	}

	var hobbies []*entity.Hobby
	if err := u.dbClient.Conn(ctx).Where("id IN ?", hobbyIDs).Find(&hobbies).Error; err != nil {
		return nil, err
	}

	return hobbies, nil
}

func (u userHobbyRepository) RegisterUserHobby(ctx context.Context, hobby *entity.Hobby) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHobbyRepository) DeleteUserHobby(ctx context.Context, userID string) error {
	//TODO implement me
	panic("implement me")
}
