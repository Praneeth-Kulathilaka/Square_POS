package routers

import (
	"Square_Pos/app/controllers"

	"github.com/gorilla/mux"
)

func SetRoute() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/order",controllers.CreateOrder).Methods("POST")
	r.HandleFunc("/order/{order_id}",controllers.GetOrder).Methods("GET")

	return r
}