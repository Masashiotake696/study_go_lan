package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; ; {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	fruits := [...]string{"Apple", "Banana", "Orange"}
	for i, s := range fruits {
		println(i, s)
	}

	text := "text"
	for _, s := range text {
		fmt.Println(s)
	}

LOOP:
	for {
		for {
			for {
				fmt.Println("開始")
				break LOOP
			}
			fmt.Println("ここは通らない")
		}
		fmt.Println("ここは通らない")
	}
	fmt.Println("完了")

L:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > 1 {
				continue L
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("ここは処理されない")
	}
}
