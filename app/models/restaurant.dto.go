package models

type Restaurant struct {
	RestaurantId int `json:"rest_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	SquareAccessKey string `json:"square_access_key"`
}