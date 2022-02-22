// Code generated by MockGen. DO NOT EDIT.
// Source: Service.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dao "stlab.itechart-group.com/go/food_delivery/courier_service/dao"
)

// MockOrderApp is a mock of OrderApp interface.
type MockOrderApp struct {
	ctrl     *gomock.Controller
	recorder *MockOrderAppMockRecorder
}

// MockOrderAppMockRecorder is the mock recorder for MockOrderApp.
type MockOrderAppMockRecorder struct {
	mock *MockOrderApp
}

// NewMockOrderApp creates a new mock instance.
func NewMockOrderApp(ctrl *gomock.Controller) *MockOrderApp {
	mock := &MockOrderApp{ctrl: ctrl}
	mock.recorder = &MockOrderAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderApp) EXPECT() *MockOrderAppMockRecorder {
	return m.recorder
}

// GetAllOrdersOfCourierService mocks base method.
func (m *MockOrderApp) GetAllOrdersOfCourierService(limit, page, idService int) ([]dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrdersOfCourierService", limit, page, idService)
	ret0, _ := ret[0].([]dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrdersOfCourierService indicates an expected call of GetAllOrdersOfCourierService.
func (mr *MockOrderAppMockRecorder) GetAllOrdersOfCourierService(limit, page, idService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrdersOfCourierService", reflect.TypeOf((*MockOrderApp)(nil).GetAllOrdersOfCourierService), limit, page, idService)
}

// GetCourierCompletedOrders mocks base method.
func (m *MockOrderApp) GetCourierCompletedOrders(limit, page, idCourier int) ([]dao.DetailedOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourierCompletedOrders", limit, page, idCourier)
	ret0, _ := ret[0].([]dao.DetailedOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourierCompletedOrders indicates an expected call of GetCourierCompletedOrders.
func (mr *MockOrderAppMockRecorder) GetCourierCompletedOrders(limit, page, idCourier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourierCompletedOrders", reflect.TypeOf((*MockOrderApp)(nil).GetCourierCompletedOrders), limit, page, idCourier)
}

// GetCourierCompletedOrdersByMonth mocks base method.
func (m *MockOrderApp) GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year int) ([]dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourierCompletedOrdersByMonth", limit, page, idService, Month, Year)
	ret0, _ := ret[0].([]dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourierCompletedOrdersByMonth indicates an expected call of GetCourierCompletedOrdersByMonth.
func (mr *MockOrderAppMockRecorder) GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourierCompletedOrdersByMonth", reflect.TypeOf((*MockOrderApp)(nil).GetCourierCompletedOrdersByMonth), limit, page, idService, Month, Year)
}
