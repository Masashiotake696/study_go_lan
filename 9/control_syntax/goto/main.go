package main

import "fmt"

func main() {
	for {
		for {
			for {
				fmt.Println("Start")
				goto DONE
			}
		}
	}
DONE:
	fmt.Println("Done")
}
