package main

func main() {
	// fmt.Println("Vim-go")
	// //interface{}
	// var x interface{}
	// x = "ss"
	// fmt.Println(x)
	// x = 1
	// fmt.Println(x)
	// x = '山' // UNICodeが入る
	// fmt.Println(x)
	// x = [...]uint8{1, 2, 3}
	// fmt.Println(x)

	// var y, z interface{}
	// y, z = 1, 2
	// z := y + z

	// a := plus(1, 5)
	// println(a)

	// hello()

	// q, r := divide(19, 7)
	// fmt.Printf("%v,%v", q, r)

	// q, _ := divide(19, 7)
	// fmt.Printf("%v", q)

	println(doSomething1())
}

// 関数(引数有り、return有り)
func plus(x, y int) int {
	return x + y
}

// 関数(引数なし、returnなし)
func hello() {
	println("hello")
	// return // 書いても良いが、基本書かない
}

// 関数(引数有り、返り値複数)
func divide(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

// 関数(引数なし、returnの型指定)
func doSomething1() (a int) {
	/**
	以下と同じ意味
	var a int
	return a
	**/
	return
}

// 関数(引数なし、returnの型定義)
func doSomething2() (a, b int) {
	/**
	以下と同じ意味
	var a,b int
	return a,b
	**/
	return
}

// 関数(引数有り、return)
func doSomething3(_, _ int) int {
	return 1
}
