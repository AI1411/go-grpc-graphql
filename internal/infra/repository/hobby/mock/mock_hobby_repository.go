// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/infra/repository/hobby/hobby_repository.go

// Package mock_hobby is a generated GoMock package.
package mock_hobby

import (
	context "context"
	reflect "reflect"

	entity "github.com/AI1411/go-grpc-graphql/internal/domain/hobby/entity"
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

// CreateHobby mocks base method.
func (m *MockRepository) CreateHobby(ctx context.Context, Hobby *entity.Hobby) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHobby", ctx, Hobby)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHobby indicates an expected call of CreateHobby.
func (mr *MockRepositoryMockRecorder) CreateHobby(ctx, Hobby interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHobby", reflect.TypeOf((*MockRepository)(nil).CreateHobby), ctx, Hobby)
}

// DeleteHobby mocks base method.
func (m *MockRepository) DeleteHobby(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHobby", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHobby indicates an expected call of DeleteHobby.
func (mr *MockRepositoryMockRecorder) DeleteHobby(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHobby", reflect.TypeOf((*MockRepository)(nil).DeleteHobby), ctx, id)
}

// GetHobby mocks base method.
func (m *MockRepository) GetHobby(ctx context.Context, id string) (*entity.Hobby, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHobby", ctx, id)
	ret0, _ := ret[0].(*entity.Hobby)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHobby indicates an expected call of GetHobby.
func (mr *MockRepositoryMockRecorder) GetHobby(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHobby", reflect.TypeOf((*MockRepository)(nil).GetHobby), ctx, id)
}