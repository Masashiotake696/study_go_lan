package main

import (
	"fmt"
	"time"
)

func receiver(ch <-chan int) {
	for {
		if i, ok := <-ch; ok {
			fmt.Println(i)
		}
	}
}

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}

		fmt.Println(name, i)
	}
	fmt.Println(name, "is done.")
}

func forReceiver(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func maltiChannelReceiver(ch1 <-chan int, ch2 <-chan int) {
	for {
		select {
		case e1 := <-ch1:
			fmt.Println("ch1: ", e1)
		case e2 := <-ch2:
			fmt.Println("ch2: ", e2)
		}
	}
}

func main() {
	ch1 := make(chan int)
	fmt.Println("cap: ", cap(ch1))
	go receiver(ch1)
	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}

	ch2 := make(chan rune, 3)
	ch2 <- 'A'
	ch2 <- 'B'
	ch2 <- 'C'
	// ch2 <- 'D'

	ch3 := make(chan string, 3)
	ch3 <- "Apple"
	ch3 <- "Banana"
	ch3 <- "Cherry"
	fmt.Println("len: ", len(ch3))
	fmt.Println("cap: ", cap(ch3))

	ch4 := make(chan int, 3)
	go receiver(ch4)
	ch4 <- 1
	ch4 <- 2
	ch4 <- 3
	ch4 <- 4
	close(ch4)

	ch5 := make(chan int, 20)
	go receive("1st goroutine", ch5)
	go receive("2nd goroutine", ch5)
	go receive("3rd goroutine", ch5)

	j := 0
	for j < 100 {
		ch5 <- j
		j++
	}
	close(ch5)

	time.Sleep(1 * time.Second)

	ch6 := make(chan int)
	go forReceiver(ch6)
	ch6 <- 1
	ch6 <- 2
	ch6 <- 3
	close(ch6)

	ch7 := make(chan int)
	ch8 := make(chan int)
	go maltiChannelReceiver(ch7, ch8)
	ch8 <- 3
	ch7 <- 2

	time.Sleep(1 * time.Second)

	ch9 := make(chan int)
	ch10 := make(chan int)
	ch11 := make(chan int)

	go func() {
		for {
			ch10 <- (<-ch9 * 2)
		}
	}()

	go func() {
		for {
			ch11 <- (<-ch10 - 1)
		}
	}()

	n := 1
LOOP:
	for {
		select {
		case ch9 <- n:
			n++
		case i := <-ch11:
			fmt.Println("received", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}
