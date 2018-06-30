package main

import (
	"fmt"
)

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

type Car struct {
	Number string
	Model  string
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s]%s", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "XXX-0123", Model: "PX512"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString)
	}
}
