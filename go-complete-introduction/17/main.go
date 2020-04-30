package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func f(s string) error {
	err := errors.New("An error has occurred.")
	return errors.Wrapf(err, "f() with %s", s)
}

func main() {
	s := "Hello"
	if err := f(s); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
