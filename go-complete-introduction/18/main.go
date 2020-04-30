package main

import (
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("bar: %w", errors.New("foo"))
	fmt.Println(err)
	fmt.Println(errors.Unwrap(err))
}
