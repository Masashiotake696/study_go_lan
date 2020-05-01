package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u1, err := url.Parse("https://www.example.com/search?a=1&b=2#top")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u1.Scheme)
	fmt.Println(u1.Host)
	fmt.Println(u1.Path)
	fmt.Println(u1.RawQuery)
	fmt.Println(u1.Fragment)
	fmt.Println(u1.IsAbs())
	fmt.Println(u1.Query())

	u2 := &url.URL{}
	u2.Scheme = "https"
	u2.Host = "google.com"
	q := u2.Query()
	q.Set("q", "Go言語")
	u2.RawQuery = q.Encode()
	fmt.Println(u2)
}
