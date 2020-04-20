package main

import (
	"fmt"
	"log"
	"os"
)

type User1 struct {
	Id    int
	Email string
}

type User2 struct {
	Id    int
	Email string
}

func (u *User2) String() string {
	return fmt.Sprintf("<%d, %s>", u.Id, u.Email)
}

func main() {
	n := 4
	fmt.Printf("%d * %d = %d\n", n, n, n*n)

	s := fmt.Sprintf("[%.2f]\n", 1.23456)
	fmt.Println(s)

	f, err := os.Create("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "|%05d|%05d|\n", 121, 33)

	fmt.Printf("%v\n", [3]int{1, 2, 3})
	fmt.Printf("%v\n", []string{"A", "B", "C"})
	fmt.Printf("%v\n", map[int]float64{1: 1.0, 2: 4.0})

	u1 := &User1{Id: 123, Email: "mail@example.com"}
	fmt.Printf("%v\n", u1)
	fmt.Printf("%+v\n", u1)
	fmt.Printf("%#v\n", u1)
	u2 := &User2{Id: 123, Email: "mail@example.com"}
	fmt.Printf("%s\n", u2)
	fmt.Printf("%v\n", u2)
	fmt.Printf("%+v\n", u2)
	fmt.Printf("%#v\n", u2)

	fmt.Print(123, 3.14, "Golang", struct{ X, Y int }{1, 2}, "\n")
	fmt.Println(123, 3.14, "Golang", struct{ X, Y int }{1, 2})
}
