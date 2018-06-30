package main

import (
	"fmt"
)

func main() {
	runDefer()
}

// deferは終了直前に動作する
func runDefer() {
	// file, err:=os.Open('/path/to/tm.txt')
	// if err!=nil {
	// 	return
	// }
	// 出力はdone, 3, 2, 1(deferは下から順に評価される)
	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	// fmt.Println("done")

	// 出力は1, 2, 3(関数自体がdeferなので出力順は記述通り)
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	fmt.Println("done")
}
