package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", GetHello).Methods("GET")

	log.Print("Starting server on port 8000.")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Display Hello
func GetHello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello")
}
