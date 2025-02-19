package models

import "github.com/golang-jwt/jwt/v5"

type Restaurant struct {
	RestaurantId    int    `json:"rest_id"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	SquareAccessKey string `json:"square_access_key"`
}

type RestaurantToken struct {
	RestaurantId      int    `json:"rest_id"`
	Square_Access_Key string `json:"square_access_key"`
	jwt.RegisteredClaims
}

type RestaurantLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}