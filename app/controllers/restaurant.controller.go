package controllers

import (
	"Square_Pos/app/db"
	"Square_Pos/app/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var restaurant models.Restaurant

	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}

	passWordInBytes, errPwd := bcrypt.GenerateFromPassword([]byte(restaurant.Password), bcrypt.DefaultCost)
	if errPwd != nil {
		http.Error(w, "Error encrypting password", http.StatusInternalServerError)
		log.Println("Error hashing password", errPwd)
		return
	}

	restaurant.RestaurantId = uuid.New().ClockSequence()
	// accesTokenInBytes, errTkn := bcrypt.GenerateFromPassword([]by)
	err = db.RegisterRestaurant(restaurant.RestaurantId,restaurant.UserName, string(passWordInBytes),restaurant.SquareAccessKey)
	if err != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		log.Println("Error inserting user",err)
		return
	}
	log.Println("Restaurant Successfully Registered")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - Successfully Registered"))

}