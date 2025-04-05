package main

import "net/http"

var members = map[string]string{
	"1": "Andy",
	"2": "Peter",
	"3": "Gabriella",
	"4": "Jordy",
}

func showContactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	http.HandleFunc("/contact", showContactPage)
	println("Starting Server...")
	http.ListenAndServe(":3000", nil)
}
