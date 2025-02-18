package main

import (
	"Square_Pos/routers"
	"fmt"
	"net/http"
)

func main() {

	r := routers.SetRoute()

	port := ":8080"
	fmt.Println("Server is running on port ",port)
	http.ListenAndServe(port, r)
}