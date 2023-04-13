package report

import (
	"context"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/report"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type GetUserReportCountUsecaseImpl interface {
	Exec(context.Context) (*grpc.GetUserReportCountResponse, error)
}

type getUserReportCountUsecaseImpl struct {
	reportRepo report.Repository
}

func NewGetUserReportCountUsecaseImpl(reportRepo report.Repository) GetUserReportCountUsecaseImpl {
	return &getUserReportCountUsecaseImpl{
		reportRepo: reportRepo,
	}
}

func (u getUserReportCountUsecaseImpl) Exec(ctx context.Context) (*grpc.GetUserReportCountResponse, error) {
	res, err := u.reportRepo.GetUserReportCount(ctx)
	if err != nil {
		return nil, err
	}

	reports := make([]*grpc.ReportCount, len(res))
	for i, r := range res {
		reports[i] = &grpc.ReportCount{
			UserId: util.NullUUIDToString(r.ReportedUserID),
			Count:  int32(r.ReportCount),
		}
	}

	return &grpc.GetUserReportCountResponse{ReportCounts: reports}, nil
}
