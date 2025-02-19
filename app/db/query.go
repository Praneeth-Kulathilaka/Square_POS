package db

import (

	// "github.com/google/uuid"
)

func RegisterRestaurant(restaurant_id int,rest_name, password, accesToken string) error {

	
	query := "INSERT INTO restaurants (rest_id, username, password, square_access_key) VALUES ($1, $2, $3, $4)"

	err := DB.QueryRow(query, restaurant_id, rest_name, password, accesToken)
	if err != nil {
		return err.Err()
	}
	return nil

}