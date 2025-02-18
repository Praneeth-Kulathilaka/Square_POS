package square

import (
	"Square_Pos/app/models"
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

// var emptyPayload models.Payload

func MakeRequest(method, endpoint string, data models.OrderRequest) ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
	url := baseUrl + endpoint

	accesToken := os.Getenv("SQUARE_ACCESS_TOKEN")
	log.Println("Access token: ",accesToken)

	var reqBody []byte
	reqBody, _ = json.Marshal(data)

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+accesToken)
	req.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in HTTP Request",err)
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("square api error: %s",body)
	}
	return body, nil

}