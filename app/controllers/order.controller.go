package controllers

import (
	"Square_Pos/app/auth"
	"Square_Pos/app/models"
	"Square_Pos/app/parser"
	"Square_Pos/app/shared"
	"Square_Pos/app/square"
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var data models.OrderRequest

	restaurant_id := r.Context().Value(auth.UserContextKey).(int)
	log.Println("Restaurant Id:",restaurant_id)

	accessToken := r.Context().Value(auth.SquareAccessTokenKey).(string)
	log.Println("Square access key: ",accessToken)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}

	response, err := square.MakeRequest(http.MethodPost, "/orders",accessToken, &data)
	if err != nil {
		http.Error(w,"Square error", http.StatusInternalServerError)
		log.Println("Error calling square function: ",err)
		return
	}

	var squareResp parser.SquareOrderResponse

	err = json.Unmarshal(response, &squareResp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newOrder := parser.ParseOrder(squareResp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newOrder)
	// shared.WriteResponse(response, w)
}

func GetOrder (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	order_id := params["order_id"]
	log.Println("Order id",order_id)
	if order_id == "" {
		http.Error(w, "Invalid order id", http.StatusBadRequest)
		return
	}

	restaurant_id := r.Context().Value(auth.UserContextKey).(int)
	log.Println("Restaurant Id:",restaurant_id)

	accessToken := r.Context().Value(auth.SquareAccessTokenKey).(string)
	log.Println("Square access key: ",accessToken)
	
	response, err := square.MakeRequest(http.MethodGet, fmt.Sprintf("/orders/%s",order_id), accessToken, nil)
	if err != nil {
		http.Error(w,"Square error", http.StatusInternalServerError)
		log.Println("Error calling square function: ",err)
		return
	}
	shared.WriteResponse(response, w)
}

func PayOrder (w http.ResponseWriter, r *http.Request){
	var payData models.PaymentRequest

	err := json.NewDecoder(r.Body).Decode(&payData)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}

	restaurant_id := r.Context().Value(auth.UserContextKey).(int)
	log.Println("Restaurant Id:",restaurant_id)

	accessToken := r.Context().Value(auth.SquareAccessTokenKey).(string)
	log.Println("Square access key: ",accessToken)

	response, err := square.MakeRequest(http.MethodPost, "/payments", accessToken, &payData)
	if err != nil {
		http.Error(w,"Square error", http.StatusInternalServerError) 
		log.Println("Error calling square function: ",err)
		return
	}
	shared.WriteResponse(response,w)
}