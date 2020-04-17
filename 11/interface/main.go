package main

import "fmt"

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

func Println(s Stringify) {
	fmt.Println(s.ToString())
}

type T struct {
	Id   int
	Name string
}

func (t *T) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.Id, t.Name)
}

type I0 interface {
	Method1() int
}

type I1 interface {
	I0
	Method2() int
}

type I3 interface {
	I1
	Method3() int
}

type TT struct {
	Id int
}

func (tt *TT) GetId() int {
	return tt.Id
}

func ShowId(id interface{ GetId() int }) {
	fmt.Println(id.GetId())
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}

	vs := []Stringify{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "XXX-123", Model: "ATS111"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	Println(&Person{Name: "Hanako", Age: 23})
	Println(&Car{Number: "AAA-1212", Model: "FFFFFss"})

	t := &T{Id: 1, Name: "Taro"}
	fmt.Println(t)

	tt := &TT{Id: 1}
	ShowId(tt)
}
