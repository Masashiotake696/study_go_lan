package main

import (
	"context"
	"fmt"
)

// func infiniteLoop(ctx context.Context) {
// 	for {
// 		fmt.Println("Help!")
// 	}
// }

// func main() {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
// 	defer cancel()

// 	go infiniteLoop(ctx)

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println(ctx.Err())
// 	}
// }

// func infiniteLoop(ctx context.Context) {
// 	innerCtx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	for {
// 		fmt.Println("Help!")

// 		select {
// 		case <-innerCtx.Done():
// 			fmt.Println("Exit from hell")
// 			return
// 		}
// 	}
// }

// func main() {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)

// 	go infiniteLoop(ctx)

// 	cancel()

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println(ctx.Err())
// 	}

// 	time.Sleep(1 * time.Second)
// }

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "hoge", 11)

	fmt.Println(ctx.Value("hoge").(int))
}
