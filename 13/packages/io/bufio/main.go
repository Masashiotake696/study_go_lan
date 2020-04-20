package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// scanner1 := bufio.NewScanner(os.Stdin)

	// for scanner1.Scan() {
	// 	fmt.Println(scanner1.Text())
	// }

	// if err := scanner1.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "エラー：", err)
	// }

	s2 := `XXXX
YYYYYYY
ZZZZZ`
	r2 := strings.NewReader(s2)
	scanner2 := bufio.NewScanner(r2)
	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
	}

	s3 := `ABC DEF
GHI JKL MNO
PQR STU UVX
YZ
`
	r3 := strings.NewReader(s3)
	scanner3 := bufio.NewScanner(r3)
	scanner3.Split(bufio.ScanWords)

	for scanner3.Scan() {
		fmt.Println(scanner3.Text())
	}

	w1 := bufio.NewWriter(os.Stdout)
	w1.WriteString("Hello, World!\n")
	w1.Flush()
}
