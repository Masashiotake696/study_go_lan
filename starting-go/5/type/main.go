package main

import "fmt"

func main() {
	n := 1
	b := byte(n)
	i64 := int64(n)
	u32 := uint32(n)
	println(b, i64, u32)

	x := [5]int{1, 2, 3}
	y := [...]int{1, 2, 3}
	z := [...]int{}
	fmt.Println(x, y, z)

	var xxx interface{}
	fmt.Println(xxx)
	xxx = 1
	fmt.Println(xxx)
	xxx = true
	fmt.Println(xxx)
	xxx = "string"
	fmt.Println(xxx)
	xxx = [...]int{1, 2, 3, 4, 5}
	fmt.Println(xxx)
}
