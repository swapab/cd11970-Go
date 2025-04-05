package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Failed to start server")
		return
	}
}
