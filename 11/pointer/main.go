package main

import "fmt"

func inc(p *int) {
	fmt.Println(p)
	fmt.Println(*p)
	*p++
}

func pow(p *[3]int) {
	i := 0
	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

func printString(s string) {
	s = "test"
	fmt.Println(s)
}

func main() {
	var i int
	p1 := &i
	fmt.Printf("%T\n", p1)
	pp1 := &p1
	fmt.Printf("%T\n", pp1)
	i = 5
	fmt.Println(*p1)
	fmt.Println(**pp1)
	*p1 = 10
	fmt.Println(i)

	j := 1
	inc(&j)
	inc(&j)
	inc(&j)
	fmt.Println(j)

	p2 := [3]int{1, 2, 3}
	pow(&p2)
	fmt.Println(p2)

	p3 := &[3]int{1, 2, 3}
	fmt.Println((*p3)[0])
	fmt.Println(p3[0])

	a := [3]string{"Apple", "Banana", "Cherry"}
	p4 := &a
	fmt.Println(a[1])
	fmt.Println(p4[1])
	p4[2] = "Grape"
	fmt.Println(a[2])
	fmt.Println(p4[2])

	p5 := &[3]int{1, 2, 3}
	fmt.Println(len(*p5))
	fmt.Println(len(p5))
	fmt.Println(cap(*p5))
	fmt.Println(cap(p5))
	fmt.Println((*p5)[0:2])
	fmt.Println(p5[0:2])

	p6 := &[3]string{"Apple", "Banana", "Cherry"}
	for i, v := range *p6 {
		fmt.Println(i, v)
	}
	for i, v := range p6 {
		fmt.Println(i, v)
	}

	c := 5
	ip := &c
	fmt.Printf("type=%T, address=%p\n", ip, ip)

	d := [3]int{1, 2, 3}
	p7 := &d[1]
	*p7 = 0
	fmt.Println(d)

	s := "Hello, World!"
	printString(s)
	fmt.Println(s)
}
