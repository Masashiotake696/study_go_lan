package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	v := struct {
		Message string `json:"message"`
	}{
		Message: "Hello",
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error: ", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
