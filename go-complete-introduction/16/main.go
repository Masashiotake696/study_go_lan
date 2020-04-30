package main

import (
	"errors"
	"fmt"

	"go.uber.org/multierr"
)

func step1() error {
	return errors.New("Error: Step1")
}

func step2() error {
	return errors.New("Error: Step2")
}

func main() {
	var rerr error
	if err := step1(); err != nil {
		rerr = multierr.Append(rerr, err)
	}
	if err := step2(); err != nil {
		rerr = multierr.Append(rerr, err)
	}

	for _, err := range multierr.Errors(rerr) {
		fmt.Println(err)
	}
}
