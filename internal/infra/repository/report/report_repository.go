package report

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/internal/domain/report/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
)

type ReportRepository interface {
	ListReport(context.Context, string) ([]*entity.Report, error)
	GetUserReportCount(context.Context, string) (int, error)
	GetReport(context.Context, string) (*entity.Report, error)
	CreateReport(context.Context, *entity.Report) error
	UpdateReportStatus(context.Context, string, string) error
}

type reportRepository struct {
	dbClient *db.Client
}

func NewReportRepository(dbClient *db.Client) ReportRepository {
	return &reportRepository{
		dbClient: dbClient,
	}
}

func (r reportRepository) ListReport(ctx context.Context, reportedUserID string) ([]*entity.Report, error) {
	var reports []*entity.Report
	if err := r.dbClient.Conn(ctx).Where("reported_user_id = ?", reportedUserID).Find(&reports).Error; err != nil {
		return nil, err
	}

	return reports, nil
}

func (r reportRepository) GetUserReportCount(ctx context.Context, reportedUserID string) (int, error) {
	var reports []*entity.Report
	if err := r.dbClient.Conn(ctx).Model(&entity.Report{}).Where("reported_user_id = ?", reportedUserID).Find(&reports).Error; err != nil {
		return 0, err
	}

	return len(reports), nil
}

func (r reportRepository) GetReport(ctx context.Context, s string) (*entity.Report, error) {
	//TODO implement me
	panic("implement me")
}

func (r reportRepository) CreateReport(ctx context.Context, e *entity.Report) error {
	//TODO implement me
	panic("implement me")
}

func (r reportRepository) UpdateReportStatus(ctx context.Context, s string, s2 string) error {
	//TODO implement me
	panic("implement me")
}
