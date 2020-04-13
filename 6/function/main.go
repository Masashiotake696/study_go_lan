package main

import "fmt"

func hello() {
	fmt.Println("Hello, World.")
	return
}

func plus(x, y int) int {
	return x + y
}

func div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func doSomething() (x, y int) {
	y = 5
	return
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}

func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	hello()

	fmt.Println(plus(1, 2))

	q, r := div(5, 3)
	fmt.Println(q, r)

	fmt.Println(doSomething())

	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(2, 3))

	plusAlias := plus
	fmt.Println(plusAlias(10, 2))

	rf := returnFunc()
	rf()
	returnFunc()()
	callFunction(returnFunc())

	fl := later()
	fmt.Println(fl("Golang"))
	fmt.Println(fl("is"))
	fmt.Println(fl("awesome!"))
	fmt.Println(fl(""))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	otherInts := integers()
	fmt.Println(otherInts())
}
