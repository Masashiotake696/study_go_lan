package main

import "fmt"

func main() {
	/*
		type T0 int
		type T1 int

		t0 := T0(5)
		i0 := int(t0)

		t1 := T1(8)
		i1 := int(t1)

		t0 = t1
	*/

	// Pointというエイリアスを構造体に貼る
	type Point struct { // 構造体はstruct
		// X int
		// Y int
		X, Y int
	}

	var pt Point
	pt.X = 10
	fmt.Println(pt)

	pt2 := Point{1, 2}
	fmt.Println(pt2)

	pt3 := Point{X: 1, Y: 2}
	fmt.Println(pt3)

	type Person struct {
		ID         uint
		name       string
		department string
	}

	p := Person{ID: 33, name: "Otake", department: "DTU"}
	fmt.Println(p)

	// 構造体を構造体に入れることもできる
	type Feed struct {
		Name   string
		Amount uint
	}

	type Animal struct {
		Name string
		// Feed Feed
		Feed // 型名を省略すると変数名と同じ型名になる
	}

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	fmt.Println(a)

	println(a.Name)
	println(a.Feed.Name)
	// println(a.Feed.Amount)
	println(a.Amount)

	type T0 struct {
		Name1 string
	}
	type T1 struct {
		T0
		Name2 string
	}
	type T2 struct {
		T1
		Name3 string
	}
	t := T2{T1: T1{T0: T0{Name1: "X"}, Name2: "Y"}, Name3: "Z"}
	println(t.Name1)
	println(t.Name2)
	println(t.Name3)

	type T struct {
		int     // 型名が未定義なので、暗黙的にint型に
		float64 // 型名が未定義なので、暗黙的にfloat64型に
		string  // 型名が未定義なので、暗黙的にstring型に
	}

	x := T{1, 3.14, "str"}
	println(x.int)
	println(x.string)
	println(x.float64)
}
