package shared

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteResponse(response []byte, w http.ResponseWriter) {
	// The function will write the status and responses from each functions
	var result map[string]interface{}
	err := json.Unmarshal(response, &result)
	if err != nil {
		log.Println("Error unmarshalling data: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}