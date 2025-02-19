package db

import (
	"database/sql"
	"fmt"
	"log"

	// "net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = sql.Open("postgres", url)

	if err != nil {
		log.Println("Error connecting DB: ",err)
		return
	}

	err = DB.Ping()
	if err != nil {
		log.Println("Failed to ping DB: ",err)
		return
	}
	log.Println("Database connected successfully")
}