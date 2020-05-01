package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, 5*time.Second)
	defer cancel()

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
