package report

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/report/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/report"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type CreateReportUsecaseImpl interface {
	Exec(context.Context, *grpc.CreateReportRequest) (*grpc.CreateReportResponse, error)
}

type createReportUsecaseImpl struct {
	reportRepo report.Repository
}

func NewCreateReportUsecaseImpl(reportRepo report.Repository) CreateReportUsecaseImpl {
	return &createReportUsecaseImpl{
		reportRepo: reportRepo,
	}
}

func (u *createReportUsecaseImpl) Exec(ctx context.Context, in *grpc.CreateReportRequest) (*grpc.CreateReportResponse, error) {
	id, err := u.reportRepo.CreateReport(ctx, &entity.Report{
		ReporterUserID: util.StringToNullUUID(in.GetReporterUserId()),
		ReportedUserID: util.StringToNullUUID(in.GetReportedUserId()),
		ReportedChatID: util.StringToNullUUID(in.GetReportedChatId()),
		Status:         entity.ReportStatusPending,
		Reason:         in.Reason,
	})
	if err != nil {
		return nil, err
	}

	return &grpc.CreateReportResponse{Id: id}, nil
}
