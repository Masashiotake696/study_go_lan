package main

import "fmt"

func sum(s ...int) (n int) {
	for _, v := range s {
		n += v
	}
	return
}

func pow_array(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
}

func pow_slice(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
}

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

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11}
	s3 := copy(s1, s2)
	fmt.Println(s3)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	s2 = []int{10, 11, 12, 13, 14, 15, 16}
	s3 = copy(s1, s2)
	fmt.Println(s3)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)

	q := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	r := q[2:4]
	fmt.Println(r)
	fmt.Println("len", len(r))
	fmt.Println("cap", cap(r))
	r = q[2:4:4]
	fmt.Println(r)
	fmt.Println("len", len(r))
	fmt.Println("cap", cap(r))
	r = q[2:4:6]
	fmt.Println(r)
	fmt.Println("len", len(r))
	fmt.Println("cap", cap(r))

	u := []string{"Apple", "Banana", "Ornage"}
	for i, v := range u {
		fmt.Printf("[%d] => %s\n", i, v)
		u = append(u, "Melon")
	}
	fmt.Println(u)

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())
	su := []int{1, 2, 3}
	fmt.Println(sum(su...))

	a1 := [3]int{1, 2, 3}
	pow_array(a1)
	fmt.Println(a1)
	a2 := []int{1, 2, 3}
	pow_slice(a2)
	fmt.Println(a2)

	var a3 [3]int
	var a4 []int
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a4 == nil)

	b1 := [5]int{1, 2, 3, 4, 5}
	b2 := b1[:2]
	fmt.Println(b2)
	fmt.Println("len:", len(b2))
	fmt.Println("cap:", cap(b2))
	b1[1] = 0
	fmt.Println(b2)

	c1 := [3]int{1, 2, 3}
	c2 := c1[:]
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println("len:", len(c2))
	fmt.Println("cap:", cap(c2))
	c2 = append(c2, 4)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println("len:", len(c2))
	fmt.Println("cap:", cap(c2))
	c1[0] = 9
	fmt.Println(c1)
	fmt.Println(c2)
}
