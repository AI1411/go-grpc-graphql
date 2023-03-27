package category

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-graphql/internal/domain/category/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type CategoryRepository interface {
	ListCategory(ctx context.Context, category *entity.CategoryCondition) ([]*entity.Category, error)
	GetCategory(ctx context.Context, id string) (*entity.Category, error)
	CreateCategory(ctx context.Context, category *entity.Category) (string, error)
	DeleteCategory(ctx context.Context, id string) error
}

type categoryRepository struct {
	dbClient *db.Client
}

func NewCategoryRepository(dbClient *db.Client) CategoryRepository {
	return &categoryRepository{
		dbClient: dbClient,
	}
}

func (c categoryRepository) ListCategory(ctx context.Context, condition *entity.CategoryCondition) ([]*entity.Category, error) {
	var categories []*entity.Category
	query := c.dbClient.Conn(ctx)
	query = repository.AddLimit(query, condition.Limit)
	query = repository.AddOffset(query, condition.Offset)
	query = repository.AddWhereLike(query, "name", condition.Name)
	if err := query.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (c categoryRepository) GetCategory(ctx context.Context, id string) (*entity.Category, error) {
	var category entity.Category
	if err := c.dbClient.Conn(ctx).
		Where("id", id).
		Preload("Hobbies").
		First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "category not found: %v", err)
		}
		return nil, err
	}

	return &category, nil
}

func (c categoryRepository) CreateCategory(ctx context.Context, category *entity.Category) (string, error) {
	if err := c.dbClient.Conn(ctx).Create(&category).Error; err != nil {
		return "", err
	}

	return util.NullUUIDToString(category.ID), nil
}

func (c categoryRepository) DeleteCategory(ctx context.Context, id string) error {
	if err := c.dbClient.Conn(ctx).Where("id", id).Delete(&entity.Category{}).Error; err != nil {
		return err
	}

	return nil
}
