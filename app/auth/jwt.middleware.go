package auth

import (
	"Square_Pos/app/models"
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my-jwt-secret-key")

type contextKey string

const (
	UserContextKey        contextKey = "restaurantID"
	SquareAccessTokenKey  contextKey = "squareAccessToken"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader,"Bearer "){
			http.Error(w, "Missing Token",http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &models.RestaurantToken{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		log.Println("Claims",claims.RestaurantId)
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Store the extracted data from the token in context
		ctx := context.WithValue(r.Context(), UserContextKey, claims.RestaurantId)
		ctx = context.WithValue(ctx,SquareAccessTokenKey, claims.Square_Access_Key)
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}