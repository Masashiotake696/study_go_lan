package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Points []*Point

func (ps Points) ToString() (str string) {
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d, %d]", p.X, p.Y)
		}
	}
	return
}

func main() {
	ps := Points{}
	ps = append(ps, &Point{X: 1, Y: 2})
	ps = append(ps, nil)
	ps = append(ps, &Point{X: 3, Y: 4})

	fmt.Println(ps.ToString())
}
