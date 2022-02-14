package model

import "github.com/Baraulia/COURIER_SERVICE/dao"

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type OrderApp interface {
	GetCourierCompletedOrders( limit,page,idCourier int) ([]dao.Order,error)
	GetAllOrdersOfCourierService(limit,page,idService int) ([]dao.Order,error)
	GetCourierCompletedOrdersByMonth(limit,page,idService,Month int) ([]dao.Order,error)
	AssigningOrderToCourier(order dao.Order) error
}

type Service struct {
	OrderApp
}
func NewService(rep *dao.Repository) *Service {
	return &Service{
		NewOrderService(*rep),
	}
}