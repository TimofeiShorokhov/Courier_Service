package db

import (
	"fmt"
)

type Courier struct {
	IdCourier        uint16 `json:"id_courier"`
	CourierName      string `json:"courier_name"`
	ReadyToGo        bool   `json:"ready_to_go"`
	PhoneNumber      string `json:"phone_number"`
	Email            string `json:"email"`
	Rating           uint16 `json:"rating"`
	Photo            string `json:"photo"`
	Surname          string `json:"surname"`
	NumberOfFailures uint16 `json:"number_of_failures"`
	Deleted          bool   `json:"deleted"`
}

type SmallInfo struct {
	IdCourier   uint16 `json:"id_courier"`
	CourierName string `json:"courier_name"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"photo"`
	Surname     string `json:"surname"`
}

func GetCouriersFromDB(Couriers *[]SmallInfo) {
	db := ConnectDB()
	defer db.Close()

	selectValue := `Select "id_courier","name", "phone_number","photo", "surname" from "couriers"`

	get, err := db.Query(selectValue)

	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.IdCourier, &courier.CourierName, &courier.PhoneNumber, &courier.Photo, &courier.Surname)
		*Couriers = append(*Couriers, courier)
	}
}

func GetOneCourierFromDB(Couriers *SmallInfo, id int) {
	db := ConnectDB()
	defer db.Close()

	selectValue := `Select id_courier,name,phone_number,photo, surname from couriers where id_courier = $1`
	get, err := db.Query(selectValue, id)

	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.IdCourier, &courier.CourierName, &courier.PhoneNumber, &courier.Photo, &courier.Surname)
		*Couriers = courier
	}
}
