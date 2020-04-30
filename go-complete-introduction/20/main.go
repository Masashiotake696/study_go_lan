package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	err := os.ErrExist
	if errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}
}
