package db

import (
	"fmt"
	"strconv"
)

type Order struct {
	IdDeliveryService uint16 `json:"id_delivery_service"`
	IdOrder           uint16 `json:"id_order"`
	IdCourier         uint16 `json:"id_courier"`
	DeliveryTime      string `json:"delivery_time"`
	CustomerAdress    string `json:"customer_adress"`
	Status            string `json:"status"`
	OrderDate         string `json:"order_date"`
}

func GetActiveOrdersFromDB(Orders *[]Order) {
	db := ConnectDB()
	defer db.Close()

	get, err := db.Query(fmt.Sprintf("Select * from delivery where status = 'active'"))

	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var order Order
		err = get.Scan(&order.IdDeliveryService, &order.IdOrder, &order.IdCourier, &order.DeliveryTime, &order.CustomerAdress, &order.Status, &order.OrderDate)
		*Orders = append(*Orders, order)
	}
}

func GetActiveOrderFromDB(Orders *Order, id string) {
	db := ConnectDB()
	defer db.Close()
	l, _ := strconv.Atoi(id)

	get, err := db.Query(fmt.Sprintf("Select * from delivery where id_order = %d AND status = 'active'", l))

	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var order Order
		err = get.Scan(&order.IdDeliveryService, &order.IdOrder, &order.IdCourier, &order.DeliveryTime, &order.CustomerAdress, &order.Status, &order.OrderDate)
		*Orders = order
	}
}
