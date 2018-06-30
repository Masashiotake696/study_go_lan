package main

import (
	"fmt"
)

// 普通の関数を変数に入れることもできる
var plusAlias = plus

func plus(x, y int) int {
	return x + y
}

func main() {

	// クロージャー(そのまま代入)
	// f := func(x, y int) int { return x + y }
	// fmt.Println(f(11, 12))
	// fmt.Printf("%T", f) // func(int, int) int

	// クロージャー(変数宣言後に代入)
	// var x func(int, int) int // 関数を代入できる変数を宣言
	// x = func(x, y int) int { return x + y }
	// fmt.Printf("%T", x)

	// 即時関数
	fmt.Println(func(x, y int) int { return x + y }(11, 12))
}
