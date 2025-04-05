package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var cityPopulations = map[string]uint32{
	"Tokyo":       37435191,
	"Delhi":       29399141,
	"Shanghai":    26317104,
	"Sao Paulo":   21846507,
	"Mexico City": 21671908,
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(cityPopulations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something Went Wring")
	}
}

func main() {
	http.HandleFunc("/", index)
	println("L4. [Handlers III] Starting Server...")
	http.ListenAndServe(":3000", nil)
}
