// Code generated by MockGen. DO NOT EDIT.
// Source: internal/worker_pool/worker_pool.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	model "timescaledb-benchmark-assignment/internal/domain/model"

	gomock "github.com/golang/mock/gomock"
)

// MockWorkersPool is a mock of WorkersPool interface.
type MockWorkersPool struct {
	ctrl     *gomock.Controller
	recorder *MockWorkersPoolMockRecorder
}

// MockWorkersPoolMockRecorder is the mock recorder for MockWorkersPool.
type MockWorkersPoolMockRecorder struct {
	mock *MockWorkersPool
}

// NewMockWorkersPool creates a new mock instance.
func NewMockWorkersPool(ctrl *gomock.Controller) *MockWorkersPool {
	mock := &MockWorkersPool{ctrl: ctrl}
	mock.recorder = &MockWorkersPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkersPool) EXPECT() *MockWorkersPoolMockRecorder {
	return m.recorder
}

// ProcessQueries mocks base method.
func (m *MockWorkersPool) ProcessQueries(queryParams []*model.QueryParam, numOfWorkers int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessQueries", queryParams, numOfWorkers)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessQueries indicates an expected call of ProcessQueries.
func (mr *MockWorkersPoolMockRecorder) ProcessQueries(queryParams, numOfWorkers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessQueries", reflect.TypeOf((*MockWorkersPool)(nil).ProcessQueries), queryParams, numOfWorkers)
}
