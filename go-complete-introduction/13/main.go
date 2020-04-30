package main

import "fmt"

type Stringer interface {
	String() string
}

type I int

func (i I) String() string {
	return "I"
}

type B bool

func (b B) String() string {
	return "B"
}

type S string

func (s S) String() string {
	return "S"
}

func F(s Stringer) {
	switch v := s.(type) {
	case I:
		fmt.Println("I", int(v))
	case B:
		fmt.Println("B", bool(v))
	case S:
		fmt.Println("S", string(v))
	}
}

func main() {
	F(I(100))
	F(B(true))
	F(S("test"))
}
