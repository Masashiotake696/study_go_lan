package main

import (
	"fmt"
)

func main() {
	// 変数に入れて使う方法
	// f := returnFunc()
	// f()

	// 変数に入れないで使う方法
	// returnFunc()()

	callFunc(func() {
		fmt.Println("hello call function")
	})
}

func callFunc(f func()) {
	f()
}

// 関数を返す関数
func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}
