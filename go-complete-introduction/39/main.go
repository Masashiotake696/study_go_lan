package main

import (
	"errors"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	var eg errgroup.Group
	eg.Go(func() error {
		err := errors.New("Error!!!")
		if err != nil {
			return err
		}
		return nil
	})
	eg.Go(func() error {
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
