package main

import (
	"fmt"
	"os"
)

func runDefer() {
	defer fmt.Println("defer")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer func() {
		fmt.Println("A")
		fmt.Println("B")
		fmt.Println("C")
	}()
	fmt.Println("done")
}

func fileProcessing(s string) {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 666)
	if err != nil {
		fmt.Println("Failed to open file test.txt")
		return
	}
	defer file.Close()

	b := []byte(s)
	_, err = file.Write(b)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	runDefer()
	fileProcessing("test dayo")
}
