package report

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/report"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type ListReportUsecaseImpl interface {
	Exec(context.Context, string) (*grpc.ListReportResponse, error)
}

type listReportUsecaseImpl struct {
	reportRepo report.ReportRepository
}

func NewListReportUsecaseImpl(reportRepo report.ReportRepository) ListReportUsecaseImpl {
	return &listReportUsecaseImpl{
		reportRepo: reportRepo,
	}
}

func (u *listReportUsecaseImpl) Exec(ctx context.Context, reportedUserID string) (*grpc.ListReportResponse, error) {
	res, err := u.reportRepo.ListReport(ctx, reportedUserID)
	if err != nil {
		return nil, err
	}

	reports := make([]*grpc.Report, len(res))
	for i, r := range res {
		reports[i] = &grpc.Report{
			Id:             util.NullUUIDToString(r.ID),
			ReporterUserId: util.NullUUIDToString(r.ReporterUserID),
			ReportedUserId: util.NullUUIDToString(r.ReportedUserID),
			ReportedChatId: util.NullUUIDToString(r.ReportedChatID),
			Status:         grpc.Report_ReportStatus(r.Status.Int()),
			Reason:         r.Reason,
			CreatedAt:      timestamppb.New(r.CreatedAt),
			UpdatedAt:      timestamppb.New(r.UpdatedAt),
		}
	}

	return &grpc.ListReportResponse{Reports: reports}, nil
}
