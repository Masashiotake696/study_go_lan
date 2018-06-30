package main

import "fmt"

func main() {
	fmt.Println("Vim-go");
	fmt.Println("1", "2", "3", "4")
	fmt.Printf("数値=%d\n", 5);
	fmt.Printf("10:%d 2:%b 8=%o 16:%x\n", 17, 17, 17, 17);
	fmt.Printf("数値=%d %d %d\n", 5, 17, 11)

	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "golang", [...]int{1, 2, 3})
	print("aiueo\n")
}