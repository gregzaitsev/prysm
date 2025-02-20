// Code generated by MockGen. DO NOT EDIT.
// Source: validator/client/beacon-api/genesis.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	beacon "github.com/prysmaticlabs/prysm/v4/beacon-chain/rpc/eth/beacon"
	httputil "github.com/prysmaticlabs/prysm/v4/network/httputil"
)

// MockGenesisProvider is a mock of GenesisProvider interface.
type MockGenesisProvider struct {
	ctrl     *gomock.Controller
	recorder *MockGenesisProviderMockRecorder
}

// MockGenesisProviderMockRecorder is the mock recorder for MockGenesisProvider.
type MockGenesisProviderMockRecorder struct {
	mock *MockGenesisProvider
}

// NewMockGenesisProvider creates a new mock instance.
func NewMockGenesisProvider(ctrl *gomock.Controller) *MockGenesisProvider {
	mock := &MockGenesisProvider{ctrl: ctrl}
	mock.recorder = &MockGenesisProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenesisProvider) EXPECT() *MockGenesisProviderMockRecorder {
	return m.recorder
}

// GetGenesis mocks base method.
func (m *MockGenesisProvider) GetGenesis(ctx context.Context) (*beacon.Genesis, *httputil.DefaultErrorJson, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenesis", ctx)
	ret0, _ := ret[0].(*beacon.Genesis)
	ret1, _ := ret[1].(*httputil.DefaultErrorJson)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGenesis indicates an expected call of GetGenesis.
func (mr *MockGenesisProviderMockRecorder) GetGenesis(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenesis", reflect.TypeOf((*MockGenesisProvider)(nil).GetGenesis), ctx)
}
