package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var pathError *os.PathError

	err := &os.PathError{Path: "dummy path", Err: errors.New("dummy error")}

	if errors.As(err, &pathError) {
		fmt.Println("Failed at path:", pathError.Path)
		fmt.Println(err)
	} else {
		fmt.Println(err)
	}
}
