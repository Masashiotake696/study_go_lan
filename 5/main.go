package main

func main() {
	/*
	var a int
	var i int

	var c, d, e string

	var (
		x, y int
		name string
	)
	*/

	var n int
	n = 5
	// n = "sss"
	print(n)

	// :=は型の定義と代入を一緒にやってくれる
	// b := true // bool
	// f := 4.31 // float
	// s := "aaaaaa" // string

	print(one())

	c := one()

	print(c)
}

func one() int {
	return 1
}