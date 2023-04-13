package report

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/repository/report"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

type GetReportUsecaseImpl interface {
	Exec(context.Context, string) (*grpc.GetReportResponse, error)
}

type getReportUsecaseImpl struct {
	reportRepo report.Repository
}

func NewGetReportUsecaseImpl(reportRepo report.Repository) GetReportUsecaseImpl {
	return &getReportUsecaseImpl{
		reportRepo: reportRepo,
	}
}

func (u *getReportUsecaseImpl) Exec(ctx context.Context, id string) (*grpc.GetReportResponse, error) {
	report, err := u.reportRepo.GetReport(ctx, id)
	if err != nil {
		return nil, err
	}

	return &grpc.GetReportResponse{Report: &grpc.Report{
		Id:             util.NullUUIDToString(report.ID),
		ReporterUserId: util.NullUUIDToString(report.ReporterUserID),
		ReportedChatId: util.NullUUIDToString(report.ReportedUserID),
		ReportedUserId: util.NullUUIDToString(report.ReportedChatID),
		Reason:         report.Reason,
		Status:         grpc.Report_ReportStatus(report.Status.Int()),
		CreatedAt:      timestamppb.New(report.CreatedAt),
		UpdatedAt:      timestamppb.New(report.UpdatedAt),
	}}, nil
}
