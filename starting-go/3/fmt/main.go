package main

import "fmt"

func main() {
	// fmt.Println
	fmt.Println("Hello, Golang!")
	fmt.Println("My", "name", "is", "Taro")

	// fmt.Printf
	fmt.Printf("%d\n", 5)
	fmt.Printf("%d %b %o %x\n", 17, 17, 17, 17)
	fmt.Printf("%v %v %v\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("%#v %#v %#v\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("%T %T %T\n", 5, "Golang", [...]int{1, 2, 3})

	// print, println (標準エラー出力)
	print("Hello, World!")
	println("Hello, World!")
	println(1, 2, 3)
}
