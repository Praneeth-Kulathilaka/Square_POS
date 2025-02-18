package controllers

import (
	"Square_Pos/app/models"
	"Square_Pos/app/square"
	"encoding/json"
	"log"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var data models.OrderRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}

	response, err := square.MakeRequest(http.MethodPost,"/orders", data)
	if err != nil {
		log.Println("Error calling square function: ",err)
		return
	}
	log.Println(response)
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