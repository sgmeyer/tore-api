package main

import (
	"encoding/json"
	"net/http"
)

func GetTores(w http.ResponseWriter, r *http.Request) {
	tores := Tores{
		Tore{URL: "https://google.com"},
		Tore{URL: "https://shawnmeyer.com"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(tores)

	if err != nil {
		panic(err)
	}
}
