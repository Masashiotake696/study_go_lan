package main

import (
	"fmt"
)

// エイリアスを貼ればレシーバーとメソッドが使える
type MyInt int

func (m MyInt) Plus(i int) int {
	return int(m) + i
}

type Point struct {
	X, Y int
}

func (p *Point) ToString() string {
	return fmt.Sprintf("[%d], [%d]", p.X, p.Y)
}

func main() {
	// println(MyInt(4).Plus(3))
	p := Point{X: 10, Y: 18}
	println(p.ToString())

	f := (*Point).ToString
	println(f(&Point{X: 7, Y: 11}))
}
