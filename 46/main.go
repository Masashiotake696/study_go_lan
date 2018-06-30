package main

import (
	"fmt"
)

type T struct {
	Id   int
	Name string
}

/*
type Stringer interface {
	String() string
}
*/
func (t *T) String() string {
	return fmt.Sprintf("<<%d,%s>>", t.Id, t.Name)
}
func main() {
	t := &T{Id: 10, Name: "Taro"}
	fmt.Println(t)
}
