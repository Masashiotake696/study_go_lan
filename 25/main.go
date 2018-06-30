package main

import (
	"fmt"
)

func main() {
	/* makeでMapを作って見る */
	// キーがintでvalueがstring
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"
	fmt.Println(m)

	/* 一発で作って見る */
	m1 := map[int]string{1: "Hello", 5: "Boy"}
	fmt.Println(m1)

	// mapの中にsliceを入れて見る
	m2 := map[int][]int{
		1: []int{1},
		3: []int{1, 2},
		4: []int{1, 2, 3},
	}
	fmt.Println(m2)

	// 上のやつを省略で書いて見る
	m3 := map[int]map[float64]string{
		1: {3.14: "円周率"},
		2: {174.2: "身長"},
		3: {64.5: "体重"},
	}
	fmt.Println(m3)

	m4 := map[int]string{1: "A", 2: "B"}
	fmt.Println(m4[1])
	fmt.Println(m4[10]) // stringの初期値である空文字が出力される

	m5 := map[int]int{1: 3, 2: 4}
	fmt.Println(m5[1])
	fmt.Println(m5[10])

	m6 := map[int]int{1: 0}
	s, ok := m6[10] // 第二引数に代入ができたかどうかのboolが入る
	if ok == true {
		// ここに処理を書いたりする
	}
	fmt.Println(s)
	fmt.Println(ok)

	// ループ処理ができる
	for k, v := range m {
		fmt.Printf("%v %v\n", k, v)
	}

	// len関数を使える
	fmt.Println(len(m))

	// キーを指定して削除ができる
	delete(m, 81)
	fmt.Println(m)
}
