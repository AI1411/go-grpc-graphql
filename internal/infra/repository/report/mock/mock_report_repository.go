// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/infra/repository/report/report_repository.go

// Package mock_report is a generated GoMock package.
package mock_report

import (
	context "context"
	reflect "reflect"

	entity "github.com/AI1411/go-grpc-graphql/internal/domain/report/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateReport mocks base method.
func (m *MockRepository) CreateReport(arg0 context.Context, arg1 *entity.Report) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReport", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReport indicates an expected call of CreateReport.
func (mr *MockRepositoryMockRecorder) CreateReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReport", reflect.TypeOf((*MockRepository)(nil).CreateReport), arg0, arg1)
}

// GetReport mocks base method.
func (m *MockRepository) GetReport(arg0 context.Context, arg1 string) (*entity.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReport", arg0, arg1)
	ret0, _ := ret[0].(*entity.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReport indicates an expected call of GetReport.
func (mr *MockRepositoryMockRecorder) GetReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReport", reflect.TypeOf((*MockRepository)(nil).GetReport), arg0, arg1)
}

// GetUserReportCount mocks base method.
func (m *MockRepository) GetUserReportCount(arg0 context.Context) ([]*entity.ReportCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserReportCount", arg0)
	ret0, _ := ret[0].([]*entity.ReportCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserReportCount indicates an expected call of GetUserReportCount.
func (mr *MockRepositoryMockRecorder) GetUserReportCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserReportCount", reflect.TypeOf((*MockRepository)(nil).GetUserReportCount), arg0)
}

// ListReport mocks base method.
func (m *MockRepository) ListReport(arg0 context.Context, arg1 string) ([]*entity.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReport", arg0, arg1)
	ret0, _ := ret[0].([]*entity.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReport indicates an expected call of ListReport.
func (mr *MockRepositoryMockRecorder) ListReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReport", reflect.TypeOf((*MockRepository)(nil).ListReport), arg0, arg1)
}

// UpdateReportStatus mocks base method.
func (m *MockRepository) UpdateReportStatus(arg0 context.Context, arg1 *entity.Report) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReportStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReportStatus indicates an expected call of UpdateReportStatus.
func (mr *MockRepositoryMockRecorder) UpdateReportStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReportStatus", reflect.TypeOf((*MockRepository)(nil).UpdateReportStatus), arg0, arg1)
}
