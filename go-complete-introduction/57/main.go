package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	v := url.Values{"key": {"Value"}, "id": {"123"}}
	resp, err := http.PostForm("http://example.com/form", v)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}
