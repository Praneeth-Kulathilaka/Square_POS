package routers

import (
	"Square_Pos/app/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetRoute() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/order",controllers.CreateOrder).Methods(http.MethodPost)
	r.HandleFunc("/order/{order_id}",controllers.GetOrder).Methods(http.MethodGet)
	r.HandleFunc("/pay",controllers.PayOrder).Methods(http.MethodPost)

	r.HandleFunc("/register",controllers.RegisterRestaurant).Methods(http.MethodPost)
	r.HandleFunc("/login",controllers.LogintoRestaurant).Methods(http.MethodPost)



	return r
}