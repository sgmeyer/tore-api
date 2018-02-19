package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	log.Print("Starting server on port 8000.")
	log.Fatal(http.ListenAndServe(":8000", router))
}
