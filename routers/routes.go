package routers

import (
	"Square_Pos/app/auth"
	"Square_Pos/app/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetRoute() *mux.Router {
	r := mux.NewRouter()
	
	r.Handle("/order",auth.AuthMiddleware(http.HandlerFunc(controllers.CreateOrder))).Methods(http.MethodPost)
	r.Handle("/order/{order_id}",auth.AuthMiddleware(http.HandlerFunc(controllers.GetOrder))).Methods(http.MethodGet)
	r.Handle("/pay",auth.AuthMiddleware(http.HandlerFunc(controllers.PayOrder))).Methods(http.MethodPost)


	r.HandleFunc("/register",controllers.RegisterRestaurant).Methods(http.MethodPost)
	r.HandleFunc("/login",controllers.LogintoRestaurant).Methods(http.MethodPost)



	return r
}