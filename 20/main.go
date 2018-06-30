package main

import (
	"fmt"
	"runtime"
)

func main() {
	// goをつけると非同期処理になる
	// go sub() // ゴルーチン
	// for i := 0; i < 100; i++ {
	// 	fmt.Println("Main Loop")
	// }

	// 無名関数でも書ける
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		fmt.Println("sub2 Loop")
	// 	}
	// }()

	fmt.Printf("NumCPU %d\n", runtime.NumCPU())
	fmt.Printf("NumGorutine %d\n", runtime.NumGoroutine())
	fmt.Printf("Version %d\n", runtime.Version())
}

func sub() {
	for i := 0; i < 100; i++ {
		fmt.Println("Sub Loop")
	}
}
