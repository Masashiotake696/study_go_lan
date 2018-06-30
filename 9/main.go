package main

import (
	"fmt"
)

func main() {
	/*
		s := "hello\nworld"
		fmt.Println(s)

		s1 := `
		Goの
		文字列リテラル
		による
		複数行の
		文字列
		`
		fmt.Println(s1)
	*/

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)

	b := [5]int{} // var b [5]int と同じ意味
	fmt.Printf("%v\n", b)

	c := [5]int{1, 2, 3}
	fmt.Printf("%v\n", c)

	d := [3]string{}
	fmt.Printf("%v\n", d)

	e := [0]int{} // 後で値をいれることもできないので、意味なし
	fmt.Printf("%v\n", e)

	f := [...]int{1, 2, 3} // 配列の数を動的に決めてくれる
	f[0] = 0
	f[2] = 2
	fmt.Println(f)

	g := [...]int8{1, 2, 3}
	g[0] = 127

	h := [...]uint8{1, 2, 3}
	h[0] = 255

	var (
		i1 [5]int
		// i2 [3]int
		i2 [5]int
	)
	i1 = i2
	fmt.Println(i1)

	j := [3]int{1, 2, 3}
	j1 := [3]int{4, 5, 6}
	j = j1
	fmt.Println(j)

	var x interface{}
	fmt.Println(x)
}
