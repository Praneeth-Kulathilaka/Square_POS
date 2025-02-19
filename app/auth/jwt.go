package auth

import (
	"Square_Pos/app/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(restaurantId int, squareAccessToken string) (string, error) {
	expirationTime := time.Now().Add(24*time.Hour)
	claims := models.RestaurantToken{
		RestaurantId: restaurantId,
		Square_Access_Key: squareAccessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	var jwtSecret = []byte("my-jwt-secret-key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println("Error signing token: ",token)
		return "", err
	}

	return tokenStr, nil
}