package controllers

import (
	"Square_Pos/app/auth"
	"Square_Pos/app/db"
	"Square_Pos/app/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRestaurant(w http.ResponseWriter, r *http.Request) {
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

func LogintoRestaurant(w http.ResponseWriter, r *http.Request){
	var restaurant models.RestaurantLogin
	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}	
	restaurant_id, password, square_access_key, errQuery := db.LoginUser(restaurant.UserName)
	if errQuery != nil {
		http.Error(w, "Invalid username", http.StatusUnauthorized)
		log.Println("Invalid username",errQuery)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(password),[]byte(restaurant.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		log.Println("Invalid password",err)
		return
	}

	token, errToken := auth.GenerateToken(restaurant_id, square_access_key)
	if errToken != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		log.Println("Error generating token",errToken)
		return
	}
	json.NewEncoder(w).Encode(token)
}