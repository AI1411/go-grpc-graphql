package report_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/domain/report/entity"
	mockReport "github.com/AI1411/go-grpc-graphql/internal/infra/repository/report/mock"
	"github.com/AI1411/go-grpc-graphql/internal/usecase/report"
	"github.com/AI1411/go-grpc-graphql/internal/util"
)

func Test_getReportUsecaseImpl_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	defer ctrl.Finish()

	mockReportRepo := mockReport.NewMockRepository(ctrl)

	getReportUsecase := report.NewGetReportUsecaseImpl(mockReportRepo)

	expected := &entity.Report{
		ID:             util.StringToNullUUID(uuid.New().String()),
		ReporterUserID: util.StringToNullUUID(uuid.New().String()),
		ReportedChatID: util.StringToNullUUID(uuid.New().String()),
		ReportedUserID: util.StringToNullUUID(uuid.New().String()),
		Status:         entity.ReportStatusPending,
		Reason:         "reason",
	}

	t.Run("success", func(t *testing.T) {
		mockReportRepo.EXPECT().GetReport(ctx, reportID).Return(expected, nil).Times(1)

		res, err := getReportUsecase.Exec(ctx, reportID)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("error", func(t *testing.T) {
		mockReportRepo.EXPECT().GetReport(ctx, reportID).Return(nil, errors.New("error creating report"))

		res, err := getReportUsecase.Exec(ctx, reportID)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
