package main

func main() {
	/*
		i := 0
		for {
			println(i)
			if i == 100 {
				break
			}
		}
		for i < 100 {
			println(i)
		}
	*/

	/*
		fruits := [3]string{"apple", "banana", "cherry"}
		// iがindex, sがvalueになる
		for i, s := range fruits {
			println(i)
			println(s)
		}
	*/

	/*
		for i, r := range "ABC" {
			fmt.Printf("%d %d ", i, r)
		}
	*/

	/*
		// 普通のswitch文(golangはbreak文がいらない)
		n := 3
		switch n {
		case 1, 2:
			println("1 or 2")
		case 3, 4:
			println("3 or 4")
		default:
			println("unknown")
		}
	*/

	/*
		// 簡易switch文(nのスコープがswitch文内になる)
		switch n := 3; n {
		case 1, 2:
			println("1 or 2")
		case 3, 4:
			println("3 or 4")
		default:
			println("unknown")
		}
	*/

	/*
		// キャスト
		var x interface{} = 3
		i, isInt := x.(int) // 第二引数にboolが入る。キャストできればtrue、できなければfalseになる
		println(i)
		println(isInt)
	*/

	/*
		// switchを使った型判定
		var x interface{} = 3
		switch v := x.(type) {
		case bool:
			print("bool")
			print(v)
		case int, uint:
			print("int, uint")
			// print(v * v)
			// fallthrough // これを書くとbreakされない
		case string:
			print("string")
			print(v)
		default:
			print("default")
		}
	*/

	// goto文
	println("A")
	goto L
	println("B")
L:
	println("C")
	for {
		for {
			for {
				goto DONE
			}
		}
	}
DONE:
}
