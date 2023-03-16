package chat

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-graphql/internal/domain/chat/entity"
	userEntity "github.com/AI1411/go-grpc-graphql/internal/domain/user/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type ChatRepository interface {
	ListChat(ctx context.Context, userID string) ([]*entity.Chat, error)
	CreateChat(ctx context.Context, chat *entity.Chat) (string, error)
	DeleteChat(ctx context.Context, chatID string) error
}

type chatRepository struct {
	dbClient *db.Client
}

func NewChatRepository(dbClient *db.Client) ChatRepository {
	return &chatRepository{
		dbClient: dbClient,
	}
}

func (c *chatRepository) ListChat(ctx context.Context, userID string) ([]*entity.Chat, error) {
	var chats []*entity.Chat
	if err := c.dbClient.Conn(ctx).Where("from_user_id", userID).Find(&chats).Error; err != nil {
		return nil, err
	}
	return chats, nil
}

func (c *chatRepository) CreateChat(ctx context.Context, chat *entity.Chat) (string, error) {
	var fromUser *userEntity.User
	if err := c.dbClient.Conn(ctx).Where("id", chat.FromUserID).First(&fromUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", status.Errorf(codes.NotFound, "from user not found")
		}
		return "", err
	}

	for _, userStatus := range userEntity.NotActiveUser {
		if fromUser.Status == userStatus {
			return "", status.Errorf(codes.FailedPrecondition, "from user is not active")
		}
	}

	var toUser *userEntity.User
	if err := c.dbClient.Conn(ctx).Where("id", chat.ToUserID).First(&toUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", status.Errorf(codes.NotFound, "to user not found")
		}
		return "", err
	}

	for _, userStatus := range userEntity.NotActiveUser {
		if toUser.Status == userStatus {
			return "", status.Errorf(codes.FailedPrecondition, "to user is not active")
		}
	}

	if err := c.dbClient.Conn(ctx).Create(chat).Error; err != nil {
		return "", err
	}

	return util.NullUUIDToString(chat.ID), nil
}

func (c *chatRepository) DeleteChat(ctx context.Context, chatID string) error {
	//TODO implement me
	panic("implement me")
}
