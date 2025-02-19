package main

import (
	"Square_Pos/app/db"
	"Square_Pos/routers"
	// "fmt"
	"log"
	"net/http"
	"os"

	// "github.com/joho/godotenv"
)

func main() {

	r := routers.SetRoute()
	
	db.InitDB()

	port := ":8080"
	log.Println("Server is starting on port ",port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Println("Server is not running",err)
		os.Exit(1)
	}
}