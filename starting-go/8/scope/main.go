package main

import (
	"fmt"

	"./foo"
)

func doSomething() (b string) {
	b = "test"
	{
		const b = "string"
	}
	return
}

func main() {
	fmt.Println(foo.A)
	fmt.Println(foo.N)
	fmt.Println(foo.DoSomething)
	fmt.Println(doSomething())
}
