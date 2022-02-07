package main

import (
	"github.com/Baraulia/COURIER_SERVICE/Controllers"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	db.ConnectDB()
	r.HandleFunc("/", Controllers.GreetingPage).Methods("GET")
	r.HandleFunc("/couriers", Controllers.GetCouriers).Methods("GET")
	r.HandleFunc("/courier/{id_courier}", Controllers.GetOneCourier).Methods("GET")
	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
