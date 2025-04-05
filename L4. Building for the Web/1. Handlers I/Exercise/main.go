package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/cities", cityList)

	fmt.Println("Starting Server...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Failed to start server")
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the world!!!")
}

func cityList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of most populous cities.\n")
	for _, city := range cities {
		fmt.Fprintf(w, "%s\n", city)
	}
}
