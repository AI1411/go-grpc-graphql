package report_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/grpc"
	"github.com/AI1411/go-grpc-graphql/internal/domain/report/entity"
	mockReport "github.com/AI1411/go-grpc-graphql/internal/infra/repository/report/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/report"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

const reportID = "123e4567-e89b-12d3-a456-426614174000"

func Test_createReportUsecaseImpl_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReportRepo := mockReport.NewMockRepository(ctrl)

	createReportUsecase := report.NewCreateReportUsecaseImpl(mockReportRepo)

	ctx := context.Background()
	newReport := &grpc.CreateReportRequest{
		ReporterUserId: uuid.New().String(),
		ReportedChatId: uuid.New().String(),
		ReportedUserId: uuid.New().String(),
		Reason:         "reason",
	}

	req := &entity.Report{
		ReporterUserID: util.StringToNullUUID(newReport.GetReporterUserId()),
		ReportedChatID: util.StringToNullUUID(newReport.GetReportedChatId()),
		ReportedUserID: util.StringToNullUUID(newReport.GetReportedUserId()),
		Status:         entity.ReportStatusPending,
		Reason:         newReport.GetReason(),
	}

	t.Run("success", func(t *testing.T) {
		mockReportRepo.EXPECT().CreateReport(ctx, req).Return(reportID, nil).Times(1)

		res, err := createReportUsecase.Exec(ctx, newReport)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("error", func(t *testing.T) {
		mockReportRepo.EXPECT().CreateReport(ctx, req).Return("", errors.New("error creating report"))

		res, err := createReportUsecase.Exec(ctx, newReport)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
