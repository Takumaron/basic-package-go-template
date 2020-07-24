// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/repository/user/repository.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	"context"
	reflect "reflect"
	entity "wantum/pkg/domain/entity"
	repository "wantum/pkg/domain/repository"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// InsertUser mocks base method
func (m *MockRepository) InsertUser(ctx context.Context, masterTx repository.MasterTx, uid, name, thumbnail string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", ctx, masterTx, uid, name, thumbnail)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUser indicates an expected call of InsertUser
func (mr *MockRepositoryMockRecorder) InsertUser(ctx, masterTx, uid, name, thumbnail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockRepository)(nil).InsertUser), ctx, masterTx, uid, name, thumbnail)
}

// SelectByPK mocks base method
func (m *MockRepository) SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByPK", ctx, masterTx, userID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByPK indicates an expected call of SelectByPK
func (mr *MockRepositoryMockRecorder) SelectByPK(ctx, masterTx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByPK", reflect.TypeOf((*MockRepository)(nil).SelectByPK), ctx, masterTx, userID)
}

// SelectByUID mocks base method
func (m *MockRepository) SelectByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByUID", ctx, masterTx, uid)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByUID indicates an expected call of SelectByUID
func (mr *MockRepositoryMockRecorder) SelectByUID(ctx, masterTx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByUID", reflect.TypeOf((*MockRepository)(nil).SelectByUID), ctx, masterTx, uid)
}

// SelectAll mocks base method
func (m *MockRepository) SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAll", ctx, masterTx)
	ret0, _ := ret[0].(entity.UserSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAll indicates an expected call of SelectAll
func (mr *MockRepositoryMockRecorder) SelectAll(ctx, masterTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockRepository)(nil).SelectAll), ctx, masterTx)
}