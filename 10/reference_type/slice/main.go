package main

import "fmt"

func main() {
	s := make([]int, 5)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	t := make([]int, 5, 10)
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(a[0:2])
	fmt.Println(a[2:])
	fmt.Println(a[:4])
	fmt.Println(a[:])
	b := "ABCDE"
	fmt.Println(b[0:2])

	x := []int{1, 2, 3}
	fmt.Println(x)
	x = append(x, 4)
	fmt.Println(x)
	x = append(x, 5, 6, 7, 8)
	fmt.Println(x)
	str := []string{"ABC", "DEF"}
	fmt.Println(str)
	str = append(str, "GHI")
	fmt.Println(str)

	y := []int{1, 2, 3}
	z := []int{4, 5, 6}
	fmt.Println(y)
	fmt.Println(z)
	yz := append(y, z...)
	fmt.Println(yz)

	p := make([]int, 0, 0)
	fmt.Printf("(A) len=%d, cap=%d\n", len(p), cap(p))
	p = append(p, 1)
	fmt.Printf("(B) len=%d, cap=%d\n", len(p), cap(p))
	p = append(p, []int{2, 3, 4}...)
	fmt.Printf("(C) len=%d, cap=%d\n", len(p), cap(p))
	p = append(p, 5)
	fmt.Printf("(D) len=%d, cap=%d\n", len(p), cap(p))
	p = append(p, 6, 7, 8, 9)
	fmt.Printf("(E) len=%d, cap=%d\n", len(p), cap(p))
}
