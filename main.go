package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// registering the /games URL to be handled by gamesHandler function
	http.HandleFunc("/games", gamesHandler)

	// Starting the service at PORT 9099
	log.Fatal(
		http.ListenAndServe(":9099", nil),
	)
}

// defining gamesHandler function
func gamesHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only POST method
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parses the json in request body and assigns it to var i
	var i any
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, "Error while reading input", http.StatusUnprocessableEntity)
		return
	}

	// Responds with the json body from request
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(i)
	if err != nil {
		http.Error(w, "Error while producing output", http.StatusInternalServerError)
		return
	}
}
