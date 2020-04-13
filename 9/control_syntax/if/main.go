package main

import "fmt"

func main() {
	x, y := 3, 5
	if n := x * y; n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}
}
