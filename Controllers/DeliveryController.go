package Controllers

import (
	"encoding/json"
	"github.com/Baraulia/COURIER_SERVICE/Models"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"net/http"
	"strconv"
)

var Orders []db.Order

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Orders = Models.GetOrders(Orders)
	json.NewEncoder(w).Encode(Orders)
}

func GetOneOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Order db.Order
	id := r.URL.Query().Get("id")
	l, _ := strconv.Atoi(id)
	Order = Models.GetOneOrder(l)
	json.NewEncoder(w).Encode(Order)
}
