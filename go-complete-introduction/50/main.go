package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Println(contentType)

	defer r.Body.Close()
	var p Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Println("Error: ", err)
	}
	fmt.Println(p)
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
