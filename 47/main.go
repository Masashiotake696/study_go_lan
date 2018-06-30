package main

import (
	"fmt"
)

/**
type I0 interface {
	Method() int
}

type I1 interface {
	I0
	Method2() int
}

type I2 interface {
	I1
	Method3() int
}
**/

type T struct {
	ID int
}

func (t *T) GetId() int {
	return t.ID
}

func main() {
	t := &T{ID: 10}
	ShowId(t)
}

func ShowId(id *T) {
	fmt.Println(id.GetId())
}
