package main

import (
	_ "math"
)

func main() {
	// fmt.Println(Pi)
	// println(AorB()) // A

	// 無限ループ
	// for {

	// }

	// ループ
	// for i := 0; i < 10; i++ {

	// }

	// if文
	// x := 0
	// if x == 1 {
	// }

	// if文
	// if true {

	// }

	// これはエラー(trueかfalseのみ)
	// if 1 {

	// }

	// xとyのスコープがこのif文だけになり、xとyのスコープを制限できる
	if x, y := 1, 2; x < y {

	}

	if _, err := doSomething(); err != nil {

	}
}

func AorB() (b string) {
	b = "A"
	{
		b := "B"
		print(b) // B
	}
	return
}

func doSomething() {

}
