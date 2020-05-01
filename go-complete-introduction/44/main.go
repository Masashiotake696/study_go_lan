package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bxcodec/faker"
)

func main() {
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, 10*time.Second)
	defer cancel()

	var count int
	ich := make(chan string)
	text := faker.Word()
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Println("Game Start!")
	fmt.Println(text)

	go func() {
		for stdin.Scan() {
			ich <- stdin.Text()
		}
	}()

	for {
		select {
		case input := <-ich:
			if text == input {
				fmt.Printf("Success!\n\n")
				count++
			} else {
				fmt.Printf("Failure...\n\n")
			}
			text = faker.Word()
			fmt.Println(text)
		case <-ctx.Done():
			fmt.Printf("Finish!!!\nScore: %d", count)
			return
		}
	}
}
