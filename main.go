package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define a struct for the JSON response
type Response struct {
	Message string `json:"message"`
}

// handler function for the POST endpoint
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.Method == http.MethodGet {
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 Internal Server Error"))
			return
		}
		defer r.Body.Close()

		// Log the request body
		fmt.Println("Received request body:", string(body))

		// Create a response instance
		response := Response{Message: "Success"}

		// Encode the response struct to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 Internal Server Error"))
			return
		}

		// Set Content-Type header and write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 Method Not Allowed"))
	}
}
func main() {

	http.HandleFunc("/", postHandler)

	fmt.Println("Server is listening on port 12000...")
	if err := http.ListenAndServe(":12000", nil); err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
