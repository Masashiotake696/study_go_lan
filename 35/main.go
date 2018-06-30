package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func swap(p *Point) { // ポインタを受け取る
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func main() {
	p := Point{X: 1, Y: 2}
	swap(&p) // 参照渡し
	fmt.Println(p)
}
