package main

import (
	"fmt"
)

type Person struct {
	Id   int
	Name string
	Area string
}

func main() {
	p := new(Person)    // &じゃなくてnewでポインタを作れる
	fmt.Println(p.Id)   // 0
	fmt.Println(p.Name) // ""
	fmt.Println(p.Area) // ""
	i := new(int)       // newは色々使える
	println(i)
}
