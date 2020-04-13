package main

import "fmt"

var S string

func init() {
	fmt.Println("init()")

	S = "A"
}

func init() {
	S += "B"
}

func init() {
	S += "C"
}

func main() {
	fmt.Println(S)
}
