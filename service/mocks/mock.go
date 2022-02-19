// Code generated by MockGen. DO NOT EDIT.
// Source: Service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	db "github.com/Baraulia/COURIER_SERVICE/db"
	gomock "github.com/golang/mock/gomock"
)

// MockDeliveryApp is a mock of DeliveryApp interface.
type MockDeliveryApp struct {
	ctrl     *gomock.Controller
	recorder *MockDeliveryAppMockRecorder
}

// MockDeliveryAppMockRecorder is the mock recorder for MockDeliveryApp.
type MockDeliveryAppMockRecorder struct {
	mock *MockDeliveryApp
}

// NewMockDeliveryApp creates a new mock instance.
func NewMockDeliveryApp(ctrl *gomock.Controller) *MockDeliveryApp {
	mock := &MockDeliveryApp{ctrl: ctrl}
	mock.recorder = &MockDeliveryAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeliveryApp) EXPECT() *MockDeliveryAppMockRecorder {
	return m.recorder
}

// ChangeOrderStatus mocks base method.
func (m *MockDeliveryApp) ChangeOrderStatus(id uint16) (uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOrderStatus", id)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeOrderStatus indicates an expected call of ChangeOrderStatus.
func (mr *MockDeliveryAppMockRecorder) ChangeOrderStatus(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderStatus", reflect.TypeOf((*MockDeliveryApp)(nil).ChangeOrderStatus), id)
}

// GetOrder mocks base method.
func (m *MockDeliveryApp) GetOrder(id int) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", id)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockDeliveryAppMockRecorder) GetOrder(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockDeliveryApp)(nil).GetOrder), id)
}

// GetOrders mocks base method.
func (m *MockDeliveryApp) GetOrders(id int) ([]db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", id)
	ret0, _ := ret[0].([]db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockDeliveryAppMockRecorder) GetOrders(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockDeliveryApp)(nil).GetOrders), id)
}

// MockCourierApp is a mock of CourierApp interface.
type MockCourierApp struct {
	ctrl     *gomock.Controller
	recorder *MockCourierAppMockRecorder
}

// MockCourierAppMockRecorder is the mock recorder for MockCourierApp.
type MockCourierAppMockRecorder struct {
	mock *MockCourierApp
}

// NewMockCourierApp creates a new mock instance.
func NewMockCourierApp(ctrl *gomock.Controller) *MockCourierApp {
	mock := &MockCourierApp{ctrl: ctrl}
	mock.recorder = &MockCourierAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCourierApp) EXPECT() *MockCourierAppMockRecorder {
	return m.recorder
}

// GetCourier mocks base method.
func (m *MockCourierApp) GetCourier(id int) (db.SmallInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourier", id)
	ret0, _ := ret[0].(db.SmallInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourier indicates an expected call of GetCourier.
func (mr *MockCourierAppMockRecorder) GetCourier(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourier", reflect.TypeOf((*MockCourierApp)(nil).GetCourier), id)
}

// GetCouriers mocks base method.
func (m *MockCourierApp) GetCouriers() ([]db.SmallInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCouriers")
	ret0, _ := ret[0].([]db.SmallInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCouriers indicates an expected call of GetCouriers.
func (mr *MockCourierAppMockRecorder) GetCouriers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCouriers", reflect.TypeOf((*MockCourierApp)(nil).GetCouriers))
}
