// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces (interfaces: Repository)

// Package livemocks is a generated GoMock package.
package livemocks

import (
	model "github.com/cshep4/premier-predictor-microservices/src/common/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
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

// GetMatchFacts mocks base method
func (m *MockRepository) GetMatchFacts(arg0 string) (*model.MatchFacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchFacts", arg0)
	ret0, _ := ret[0].(*model.MatchFacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchFacts indicates an expected call of GetMatchFacts
func (mr *MockRepositoryMockRecorder) GetMatchFacts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchFacts", reflect.TypeOf((*MockRepository)(nil).GetMatchFacts), arg0)
}

// GetUpcomingMatches mocks base method
func (m *MockRepository) GetUpcomingMatches() ([]*model.MatchFacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpcomingMatches")
	ret0, _ := ret[0].([]*model.MatchFacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpcomingMatches indicates an expected call of GetUpcomingMatches
func (mr *MockRepositoryMockRecorder) GetUpcomingMatches() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpcomingMatches", reflect.TypeOf((*MockRepository)(nil).GetUpcomingMatches))
}