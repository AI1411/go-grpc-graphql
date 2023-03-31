package server

import (
	"context"

	"go.uber.org/zap"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/infra/db"
	reportRepo "github.com/AI1411/go-grpc-graphql/internal/infra/repository/report"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/report"
)

type ReportServer struct {
	grpc.UnimplementedReportServiceServer
	dbClient   *db.Client
	zapLogger  *zap.Logger
	reportRepo reportRepo.ReportRepository
}

func NewReportServer(
	dbClient *db.Client,
	zapLogger *zap.Logger,
	reportRepo reportRepo.ReportRepository,
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
