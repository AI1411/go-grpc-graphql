package report

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/internal/domain/report/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

const reportThreshold = 2

type Repository interface {
	ListReport(context.Context, string) ([]*entity.Report, error)
	GetUserReportCount(context.Context) ([]*entity.ReportCount, error)
	GetReport(context.Context, string) (*entity.Report, error)
	CreateReport(context.Context, *entity.Report) (string, error)
	UpdateReportStatus(context.Context, *entity.Report) error
}

type reportRepository struct {
	dbClient *db.Client
}

func NewReportRepository(dbClient *db.Client) Repository {
	return &reportRepository{
		dbClient: dbClient,
	}
}

func (r *reportRepository) ListReport(ctx context.Context, reportedUserID string) ([]*entity.Report, error) {
	var reports []*entity.Report
	if err := r.dbClient.Conn(ctx).Where("reported_user_id = ?", reportedUserID).Find(&reports).Error; err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *reportRepository) GetUserReportCount(ctx context.Context) ([]*entity.ReportCount, error) {
	var reports []*entity.ReportCount
	if err := r.dbClient.Conn(ctx).
		Model(&entity.Report{}).
		Select("reported_user_id, COUNT(*) AS report_count").
		Group("reported_user_id").
		Having("COUNT(*) >= ?", reportThreshold).
		Order("report_count DESC").
		Find(&reports).Error; err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *reportRepository) GetReport(ctx context.Context, id string) (*entity.Report, error) {
	var report entity.Report
	if err := r.dbClient.Conn(ctx).Where("id = ?", id).First(&report).Error; err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *reportRepository) CreateReport(ctx context.Context, report *entity.Report) (string, error) {
	if err := r.dbClient.Conn(ctx).Create(report).Error; err != nil {
		return "", err
	}

	return util.NullUUIDToString(report.ID), nil
}

func (r *reportRepository) UpdateReportStatus(ctx context.Context, report *entity.Report) error {
	if err := r.dbClient.Conn(ctx).Model(&entity.Report{}).Where("id = ?", report.ID).
		Select("Status").Updates(&report).Error; err != nil {
		return err
	}

	return nil
}
