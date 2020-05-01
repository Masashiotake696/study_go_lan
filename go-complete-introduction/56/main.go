package main

import (
	"fmt"
	_ "image/jpeg"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("./cafe.jpg")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://example.com/upload", "image/jpeg", file)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}
