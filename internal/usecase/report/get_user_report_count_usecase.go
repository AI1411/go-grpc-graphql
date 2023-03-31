package report

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/report"
)

type GetUserReportCountUsecaseImpl interface {
	Exec(context.Context, string) (int, error)
}

type getUserReportCountUsecaseImpl struct {
	reportRepo report.ReportRepository
}

func NewGetUserReportCountUsecaseImpl(reportRepo report.ReportRepository) GetUserReportCountUsecaseImpl {
	return &getUserReportCountUsecaseImpl{
		reportRepo: reportRepo,
	}
}

func (u getUserReportCountUsecaseImpl) Exec(ctx context.Context, reportedUserID string) (int, error) {
	res, err := u.reportRepo.GetUserReportCount(ctx, reportedUserID)
	if err != nil {
		return 0, err
	}
	return res, nil
}
