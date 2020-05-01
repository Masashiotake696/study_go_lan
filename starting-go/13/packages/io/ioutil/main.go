package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f1, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	bs1, err := ioutil.ReadAll(f1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bs1)
	fmt.Println(string(bs1))

	bs2, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bs2)
	fmt.Println(string(bs2))

	f3, err := ioutil.TempFile(os.TempDir(), "foo")
	if err != nil {
		log.Fatal(err)
	}
	defer f3.Close()
	f3.WriteString("Hello, World!\n")
	fmt.Println(f3.Name())
}
