// Code generated by MockGen. DO NOT EDIT.
// Source: ../work_planner/pkg/repository/internal/ports/shift_port.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	context "context"
	reflect "reflect"

	domain "github.com/akhi19/work_planner/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockIShift is a mock of IShift interface.
type MockIShift struct {
	ctrl     *gomock.Controller
	recorder *MockIShiftMockRecorder
}

// MockIShiftMockRecorder is the mock recorder for MockIShift.
type MockIShiftMockRecorder struct {
	mock *MockIShift
}

// NewMockIShift creates a new mock instance.
func NewMockIShift(ctrl *gomock.Controller) *MockIShift {
	mock := &MockIShift{ctrl: ctrl}
	mock.recorder = &MockIShiftMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIShift) EXPECT() *MockIShiftMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockIShift) Delete(ctx context.Context, id domain.SqlID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIShiftMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIShift)(nil).Delete), ctx, id)
}

// GetShiftByID mocks base method.
func (m *MockIShift) GetShiftByID(ctx context.Context, shiftID domain.SqlID) (*domain.ShiftDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShiftByID", ctx, shiftID)
	ret0, _ := ret[0].(*domain.ShiftDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShiftByID indicates an expected call of GetShiftByID.
func (mr *MockIShiftMockRecorder) GetShiftByID(ctx, shiftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShiftByID", reflect.TypeOf((*MockIShift)(nil).GetShiftByID), ctx, shiftID)
}

// GetShiftDetails mocks base method.
func (m *MockIShift) GetShiftDetails(ctx context.Context) ([]domain.ShiftDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShiftDetails", ctx)
	ret0, _ := ret[0].([]domain.ShiftDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShiftDetails indicates an expected call of GetShiftDetails.
func (mr *MockIShiftMockRecorder) GetShiftDetails(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShiftDetails", reflect.TypeOf((*MockIShift)(nil).GetShiftDetails), ctx)
}

// Insert mocks base method.
func (m *MockIShift) Insert(ctx context.Context, shiftDTO domain.ShiftDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, shiftDTO)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockIShiftMockRecorder) Insert(ctx, shiftDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIShift)(nil).Insert), ctx, shiftDTO)
}