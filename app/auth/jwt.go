package auth

import (
	"Square_Pos/app/models"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(restaurantId int, squareAccessToken string) (string, error) {
	err:= godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	expirationTime := time.Now().Add(24*time.Hour)
	claims := models.RestaurantToken{
		RestaurantId: restaurantId,
		Square_Access_Key: squareAccessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	secret := os.Getenv("JWT_SECRET_KEY")
	var jwtSecret = []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println("Error signing token: ",token)
		return "", err
	}

	return tokenStr, nil
}