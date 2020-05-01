package main

import "fmt"

func judgmentEvenAndOdd(i int) string {
	if i%2 == 0 {
		return fmt.Sprintf("-偶数\n")
	}

	return fmt.Sprintf("-奇数\n")
}

func main() {
	for i := 1; i <= 100; i++ {
		print(i, judgmentEvenAndOdd(i))
	}
}
