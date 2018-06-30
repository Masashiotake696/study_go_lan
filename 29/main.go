package main

import "fmt"

func main() {
	/*
		// チャンネルの生成
		ch1 := make(chan int, 1)
		ch2 := make(chan int, 1)
		ch3 := make(chan int, 1)

		// ch1とch2に送信
		ch1 <- 1
		ch2 <- 2

		// case文の内容がランダムに実行される
		select {
		case <-ch1: // ch1から受信
			fmt.Println("recieve from ch1")
		case <-ch2: // ch2から受信
			fmt.Println("recieve from ch2")
		case ch3 <- 3: // ch3に送信
			fmt.Println("send to ch3")
		default: // ここは実行されない
			fmt.Println("Never here")
		}
	*/

	// チャンネルの生成
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// ch1から受信した整数を2倍してch2へ送信
	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()

	// ch2から受信した整数を1減算してch3へ送信
	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()

	n := 1
LOOP:
	for {
		// case文の内容がランダムに実行される
		select {
		case ch1 <- n: // 整数を1から増分させながらch1へ送信
			n++
		case i := <-ch3: // ch3から受信した整数を出力
			fmt.Println("recieved ch3", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}
