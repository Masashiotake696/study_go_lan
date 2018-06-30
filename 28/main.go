package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 20)
	go receve("1st goroutine", ch)
	go receve("2nd goroutine", ch)
	go receve("3rd goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)
	time.Sleep(3 * time.Second)
}

func receve(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "is done")
}
