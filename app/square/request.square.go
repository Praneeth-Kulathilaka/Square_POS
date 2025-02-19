package square

import (
	// "Square_Pos/app/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	// "io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const baseUrl = "https://connect.squareupsandbox.com/v2"


func MakeRequest(method, endpoint string, data interface{}) ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
	url := baseUrl + endpoint
	log.Println("URL: ",url)

	accesToken := os.Getenv("SQUARE_ACCESS_TOKEN")
	log.Println("Access token: ",accesToken)

	var reqBody []byte
	var req *http.Request
	if method == http.MethodPost {
		reqBody, _ = json.Marshal(data)	
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	} else if method == http.MethodGet {
		req, _ = http.NewRequest(method, url, nil)
	}

	req.Header.Set("Square-Version", "2025-01-23")
	req.Header.Set("Authorization", "Bearer "+accesToken)
	req.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	
	if err != nil {
		log.Println("Error in HTTP Request",err)
		return nil, err
	}
	// log.Println("Response",resp)
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("square api error: %s",body)
	}
	return body, nil

}