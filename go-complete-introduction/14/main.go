package main

import (
	"fmt"
	"log"
)

type Stringer interface {
	String() string
}

func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}

	return nil, MyError{Message: "CastError"}
}

type MyError struct {
	Message string
}

func (me MyError) Error() string {
	return me.Message
}

type S string

func (s S) String() string {
	return "S"
}

type I int

func Check(v interface{}) {
	if r, err := ToStringer(v); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(r.String())
	}
}

func main() {
	Check(S("test"))
	Check(I(100))
}
