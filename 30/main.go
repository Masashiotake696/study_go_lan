package main

import (
	"fmt"
)

func main() {
	// ポインタ
	var p *int
	fmt.Println(p == nil)

	var (
		_ *[]string
		_ *map[int]rune
		_ *chan int
	)

	var i int
	p1 := &i
	fmt.Printf("%T\n", p1)
	p2 := &p1
	fmt.Printf("%T\n", p2)

	var i3 int
	p3 := &i3 // i3 === *p3
	fmt.Println(*p3)
	*p3 = 10
	fmt.Println(i3)

	i4 := 1
	increment(&i4)
	increment(&i4)
	increment(&i4)
	increment(&i4)
	fmt.Println(i4)
}

func increment(p *int) {
	*p++
}
