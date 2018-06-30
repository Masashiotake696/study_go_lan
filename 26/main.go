package main

import (
	"fmt"
)

func main() {
	// チャネル型(非同期の関数に処理の途中で値を渡す型)
	// var (
	// 	// 送受信は送信と受信に値を代入できる(逆は不可)
	// 	ch1 chan int   // 送受信
	// 	ch2 <-chan int // 送信
	// 	ch3 chan<- int // 受信
	// )

	// チャネル型を作る
	ch := make(chan int)
	go receiver(ch)
	i := 0
	for i < 10000 {
		ch <- i
		i++
	}
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}
