package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var members = map[string]string{
	"1": "Andy",
	"2": "Peter",
	"3": "Gabriella",
	"4": "Jordy",
}

func showMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(members)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something Went Wring")
		return
	}
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	member := members[id]
	if len(member) <= 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		delete(members, id)
	}
	err := json.NewEncoder(w).Encode(members)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something Went Wrong")
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/members", showMembers).Methods("GET")
	router.HandleFunc("/member/{id}", deleteMember).Methods("DELETE")
	println("Starting Server...")
	http.ListenAndServe(":3000", router)
}
