package main

import "fmt"

const X = 1

const (
	A = 1
	B = 2
	C = 3
)

const (
	a = 1
	b
	c
	d = "あ"
	e
)

const ONE = 1

const (
	x = iota
	y
	z
)

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}

func main() {
	fmt.Println(X)
	fmt.Println(A, B, C)

	one, two := one()
	fmt.Println(one, two)

	fmt.Println(a, b, c, d, e)

	fmt.Println(x, y, z)
}
