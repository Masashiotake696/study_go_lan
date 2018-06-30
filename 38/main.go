package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type FloatPoint struct {
	X, Y float64
}

func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

func (p *FloatPoint) Render() {
	fmt.Printf("<%f, %f>\n", p.X, p.Y)
}

func main() {
	p := &Point{X: 5, Y: 13}
	p.Render()
	p1 := &FloatPoint{X: 5, Y: 13}
	p1.Render()
}
