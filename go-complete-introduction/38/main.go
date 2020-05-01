package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("First")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Second")
		wg.Done()
	}()

	wg.Wait()
}
