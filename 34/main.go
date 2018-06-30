package main

import (
	"fmt"
)

func main() {
	type Base struct {
		Id    int
		Owner string
	}
	type A struct {
		Base
		Name string
		Area string
	}
	type B struct {
		Base
		Title  string
		Bodies []string
	}
	a := A{
		Base: Base{17, "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}
	b := B{
		Base:   Base{71, "Jiro"},
		Title:  "no title",
		Bodies: []string{"A", "B"},
	}

	println(a.Id)
	println(a.Owner)
	println(a.Name)
	println(b.Id)
	println(b.Owner)
	println(b.Bodies)

	/*
		struct {
			T1	  // T1型
			*T2   // T2型
			P.T3  // T3型
			*P.T4 // T4型
		}
	*/

	/*
		type T0 struct {
			T1
		}

		type T1 struct {
			T0
		}
	*/

	s := struct{ X, Y int }{X: 1, Y: 23}
	showStruct(s)

	type Point struct {
		X, Y int
	}
	p := Point{X: 3, Y: 10}
	showStruct(p)
}

func showStruct(s struct{ X, Y int }) {
	fmt.Println(s)
}
