package main

func main() {
	ch := make(chan int, 10)
	ch <- 5
	i := <-ch
	println(i)
	// チャンネルを閉じる
	close(ch)
}
