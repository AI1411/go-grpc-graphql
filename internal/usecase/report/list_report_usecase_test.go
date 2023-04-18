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

const userID = "c0a80101-0000-0000-0000-000000000001"

func Test_listReportUsecaseImpl_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReportRepo := mockReport.NewMockRepository(ctrl)

	listReportUsecase := report.NewListReportUsecaseImpl(mockReportRepo)

	ctx := context.Background()

	expected := []*entity.Report{
		{
			ID:             util.StringToNullUUID(uuid.New().String()),
			ReporterUserID: util.StringToNullUUID(uuid.New().String()),
			ReportedChatID: util.StringToNullUUID(uuid.New().String()),
			ReportedUserID: util.StringToNullUUID(uuid.New().String()),
			Status:         entity.ReportStatusPending,
			Reason:         "reason",
		},
		{
			ID:             util.StringToNullUUID(uuid.New().String()),
			ReporterUserID: util.StringToNullUUID(uuid.New().String()),
			ReportedChatID: util.StringToNullUUID(uuid.New().String()),
			ReportedUserID: util.StringToNullUUID(uuid.New().String()),
			Status:         entity.ReportStatusPending,
			Reason:         "reason",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockReportRepo.EXPECT().ListReport(ctx, gomock.Any()).Return(expected, nil).Times(1)

		res, err := listReportUsecase.Exec(ctx, userID)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("error", func(t *testing.T) {
		mockReportRepo.EXPECT().ListReport(ctx, gomock.Any()).Return(nil, errors.New("list report error report"))

		res, err := listReportUsecase.Exec(ctx, userID)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
