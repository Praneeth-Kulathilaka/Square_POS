package controllers

import (
	"Square_Pos/app/models"
	"Square_Pos/app/square"
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var data models.OrderRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}

	response, err := square.MakeRequest(http.MethodPost, "/orders", &data)
	if err != nil {
		log.Println("Error calling square function: ",err)
		return
	}
	// log.Println(response)

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		log.Println("Error unmarshalling data: ",err)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetOrder (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	order_id := params["order_id"]
	log.Println("Order id",order_id)
	if order_id == "" {
		http.Error(w, "Invalid order id", http.StatusBadRequest)
		return
	}
	
	response, err := square.MakeRequest(http.MethodGet, fmt.Sprintf("/orders/%s",order_id), nil)
	if err != nil {
		log.Println("Error calling square function: ",err)
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		log.Println("Error unmarshalling data: ",err)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func PayOrder (w http.ResponseWriter, r *http.Request){
	var payData models.PaymentRequest

	err := json.NewDecoder(r.Body).Decode(&payData)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}
	response, err := square.MakeRequest(http.MethodPost, "/payments", &payData)
	if err != nil {
		log.Println("Error calling square function: ",err)
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		log.Println("Error unmarshalling data: ",err)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}