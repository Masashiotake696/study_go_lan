package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	n  bool
	cl int
)

func init() {
	flag.BoolVar(&n, "n", false, "行数の表示")
}

func getFileContent(fn string) (s string) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if n {
			cl++
			s += fmt.Sprintf("%d: %s\n", cl, scanner.Text())
		} else {
			s += fmt.Sprintf("%s", scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	flag.Parse()
	args := flag.Args()

	var s string
	for _, v := range args {
		s += getFileContent(v)
	}

	fmt.Fprint(os.Stdout, s)
}
