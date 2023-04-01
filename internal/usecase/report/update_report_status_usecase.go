package report

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/report/entity"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/report"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type UpdateReportStatus interface {
	Exec(context.Context, *grpc.UpdateReportStatusRequest) error
}

type updateReportStatusImpl struct {
	reportRepo report.ReportRepository
}

func NewUpdateReportStatusImpl(reportRepo report.ReportRepository) UpdateReportStatus {
	return &updateReportStatusImpl{
		reportRepo: reportRepo,
	}
}

func (u *updateReportStatusImpl) Exec(ctx context.Context, in *grpc.UpdateReportStatusRequest) error {
	err := u.reportRepo.UpdateReportStatus(ctx, &entity.Report{
		ID:     util.StringToNullUUID(in.GetId()),
		Status: entity.ReportStatusName[int32(in.GetStatus())],
	})
	if err != nil {
		return err
	}

	return nil
}
