package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Set(x, y int) {
	p.X = x
	p.Y = y
}

func main() {
	p1 := Point{}
	p1.Set(1, 2)
	fmt.Println(p1)

	ps := make([]Point, 5) // 5個分でスライス
	for _, p := range ps {
		fmt.Println(p.X, p.Y)
	}
}
