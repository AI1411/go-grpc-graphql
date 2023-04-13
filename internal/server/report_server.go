package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	reportRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/report"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/report"
)

type ReportServer struct {
	grpc.UnimplementedReportServiceServer
	dbClient   *db.Client
	zapLogger  *zap.Logger
	reportRepo reportRepo.Repository
}

func NewReportServer(
	dbClient *db.Client,
	zapLogger *zap.Logger,
	reportRepo reportRepo.Repository,
) *ReportServer {
	return &ReportServer{
		dbClient:   dbClient,
		zapLogger:  zapLogger,
		reportRepo: reportRepo,
	}
}

func (s *ReportServer) ListReport(ctx context.Context, in *grpc.ListReportRequest) (*grpc.ListReportResponse, error) {
	usecase := report.NewListReportUsecaseImpl(s.reportRepo)
	res, err := usecase.Exec(ctx, in.GetReportedUserId())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ReportServer) GetUserReportCount(ctx context.Context, _ *grpc.GetUserReportCountRequest) (*grpc.GetUserReportCountResponse, error) {
	usecase := report.NewGetUserReportCountUsecaseImpl(s.reportRepo)
	res, err := usecase.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ReportServer) GetReport(ctx context.Context, in *grpc.GetReportRequest) (*grpc.GetReportResponse, error) {
	usecase := report.NewGetReportUsecaseImpl(s.reportRepo)
	res, err := usecase.Exec(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ReportServer) CreateReport(ctx context.Context, in *grpc.CreateReportRequest) (*grpc.CreateReportResponse, error) {
	usecase := report.NewCreateReportUsecaseImpl(s.reportRepo)
	res, err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ReportServer) UpdateReportStatus(ctx context.Context, in *grpc.UpdateReportStatusRequest) (*emptypb.Empty, error) {
	usecase := report.NewUpdateReportStatusImpl(s.reportRepo)
	err := usecase.Exec(ctx, in)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
