package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type errWriter struct {
	w   io.Writer
	err error
}

func (e *errWriter) Write(p []byte) {
	if e.err != nil {
		return
	}
	_, e.err = e.w.Write(p)
}

func (e *errWriter) Err() error {
	return e.err
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	ew := &errWriter{w: f}
	ew.Write(([]byte)("test"))
	ew.Write(([]byte)("test"))
	ew.Write(([]byte)("test"))
	if ew.Err() != nil {
		fmt.Println(ew.Err())
	}
}
