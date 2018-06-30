package main

func main() {
	/**
	println(sum(1, 2, 3))
	println(sum(1, 2, 3, 4, 5))
	println(sum())

	s := []int{1, 2, 3}
	// slice型を渡すときは...を使う
	println(sum(s...))
	**/

	/*
		// 値渡し
		a := [...]int{1, 2, 3}
		pow(a)
		fmt.Println(a)

		// 参照渡し
		a1 := []int{1, 2, 3}
		pow2(a1)
		fmt.Println(a1)
	*/

	/*
		var (
			b  [3]int // [0 0 0]
			b1 []int  // [] (nil)
		)
		fmt.Println(b)
		fmt.Println(b1)
	*/

	/*
		// 配列
		c := [5]int{1, 2, 3, 4, 5}
		// slice型になる(cを参照する)
		s := c[0:2]
		fmt.Println(c)
		fmt.Println(s)
		c[1] = 10
		fmt.Println(s)
		s = append(s, 15, 1, 1, 1, 1)
		c[1] = 20
		// [1 10 15 1 1 1 1]
		// 参照している配列の容量を超えた時に参照渡しから値渡しに変わる
		fmt.Println(s)
	*/
}

// sにはslice型が入る
func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

// 値渡し
func pow(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

// 参照渡し
func pow2(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}
