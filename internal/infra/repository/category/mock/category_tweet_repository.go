// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/infra/repository/category/category_repository.go

// Package mock_category is a generated GoMock package.
package mock_category

import (
	context "context"
	reflect "reflect"

	entity "github.com/AI1411/go-grpc-graphql/internal/domain/category/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// CreateCategory mocks base method.
func (m *MockCategoryRepository) CreateCategory(ctx context.Context, category *entity.Category) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", ctx, category)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockCategoryRepositoryMockRecorder) CreateCategory(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockCategoryRepository)(nil).CreateCategory), ctx, category)
}

// DeleteCategory mocks base method.
func (m *MockCategoryRepository) DeleteCategory(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockCategoryRepositoryMockRecorder) DeleteCategory(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockCategoryRepository)(nil).DeleteCategory), ctx, id)
}

// GetCategory mocks base method.
func (m *MockCategoryRepository) GetCategory(ctx context.Context, id string) (*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", ctx, id)
	ret0, _ := ret[0].(*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockCategoryRepositoryMockRecorder) GetCategory(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategory), ctx, id)
}

// ListCategory mocks base method.
func (m *MockCategoryRepository) ListCategory(ctx context.Context, category *entity.CategoryCondition) ([]*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCategory", ctx, category)
	ret0, _ := ret[0].([]*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCategory indicates an expected call of ListCategory.
func (mr *MockCategoryRepositoryMockRecorder) ListCategory(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCategory", reflect.TypeOf((*MockCategoryRepository)(nil).ListCategory), ctx, category)
}