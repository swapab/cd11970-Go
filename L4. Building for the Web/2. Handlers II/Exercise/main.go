package main

import (
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
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	fmt.Fprintf(w, "<h1>Population by City</h1>")
	for city, population := range cityPopulations {
		fmt.Fprintf(w, "<h2>%s : %d</h2>", city, population)
	}
}

func main() {
	http.HandleFunc("/", index)
	println("L4. [Handlers II] Starting Server...")
	http.ListenAndServe(":3000", nil)
}
