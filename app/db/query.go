package db

// import "net/http"

// "github.com/google/uuid"

func RegisterRestaurant(restaurant_id int,rest_name, password, accesToken string) error {	
	query := "INSERT INTO restaurants (rest_id, username, password, square_access_key) VALUES ($1, $2, $3, $4)"

	err := DB.QueryRow(query, restaurant_id, rest_name, password, accesToken)
	if err != nil {
		return err.Err()
	}
	return nil
}

func LoginUser(username string) (int, string, string, error){
	query := "SELECT rest_id, password, square_access_key FROM restaurants WHERE username = $1"

	var restaurant_id int
	var password, square_access_key string
	
	err := DB.QueryRow(query,username).Scan(&restaurant_id, &password, &square_access_key)
	if err != nil {
		return 0, "", "", err
	}
	return restaurant_id, password, square_access_key, nil
}