package main

import "fmt"

func main() {
	/**
	参照型は種類が3つある
		slice
		map
		channel
	参照型のインスタンスを作成するときに使う関数
		make
	**/

	/*
		配列の中身がないのがslice型。
		slice型は可変
	*/
	var a []int
	fmt.Println(a)

	/*
		makeでも作れる
		makeで作る場合は長さを指定する(長さを指定しても可変長)
	*/
	s := make([]int, 10)
	fmt.Println(s)
	// 長さを取得するときはlen関数を使う
	fmt.Println(len(s))

	/*
		簡易slice文
	*/
	// 配列
	a1 := [5]int{1, 2, 3, 4, 5}
	// 配列の0番目から2個分を格納(slice型)
	a2 := a1[0:2]
	fmt.Println(a2)
	fmt.Println(a1[len(a1)-2])

	a3 := "ABCDE"[1:2]
	fmt.Println(a3)

	a4 := "あいうえお"[3:9]
	fmt.Println(a4)

	a5 := []int{1, 2, 3}
	a5 = append(a5, 4)
	a5 = append(a5, 5, 6, 7)
	fmt.Println(a5)

	a6 := []int{1, 2, 3}
	a7 := []int{2, 3, 4}
	// ...をつけるとslice型を「,」つなぎに分解してくれる
	a8 := append(a6, a7...)
	fmt.Println(a8)
}
