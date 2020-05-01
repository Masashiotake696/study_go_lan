package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch s := rand.Intn(6) + 1; s {
	case 6:
		fmt.Fprint(w, "大吉")
	case 5, 4:
		fmt.Fprint(w, "中吉")
	case 3, 2:
		fmt.Fprint(w, "吉")
	default:
		fmt.Fprint(w, "凶")
	}
}

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
