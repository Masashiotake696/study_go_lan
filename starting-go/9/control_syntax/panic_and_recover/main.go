package main

import "fmt"

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Println("panic: unknown")
			}
		}
	}()
	panic(src)
}

func main() {
	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
}
