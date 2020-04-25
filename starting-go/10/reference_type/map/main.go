package main

import "fmt"

func main() {
	m1 := make(map[int]string)
	m1[1] = "US"
	m1[81] = "JN"
	m1[86] = "CN"
	fmt.Println(m1)

	m2 := make(map[string]string)
	m2["Yamada"] = "Taro"
	m2["Sato"] = "Hanako"
	m2["Yamada"] = "Jiro"
	fmt.Println(m2)

	m3 := map[int]string{
		1: "Taro",
		2: "Hanako",
		3: "Jiro",
	}
	fmt.Println(m3)

	m4 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}
	fmt.Println(m4)

	m5 := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	fmt.Println(m5)

	m6 := map[int]map[int]string{
		1: {
			1: "Taro",
			2: "Jiro",
			3: "Saburo",
		},
	}
	fmt.Println(m6)

	n1 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	n2 := n1[1]
	n3 := n1[5]
	fmt.Println(n2, n3)

	n4 := map[int]int{}
	fmt.Println(n4[0], n4[4])

	n5, ok := n1[1]
	fmt.Println(n5, ok)
	n6, ok := n1[3]
	fmt.Println(n6, ok)
	n7, ok := n1[5]
	fmt.Println(n7, ok)

	if _, ok := n1[2]; ok {
		fmt.Println("OK")
	}

	n8 := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	n9 := n8[1]
	if n9 != nil {
		fmt.Println("Bad Example")
	}
	if _, ok := n8[1]; ok {
		fmt.Println("Good Example")
	}

	a1 := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Orange",
	}
	for k, v := range a1 {
		fmt.Printf("%d => %s\n", k, v)
	}

	a2 := map[string]int{
		"Apple":  88,
		"Banana": 107,
		"Cherry": 46,
	}
	a2["Grape"] = 66
	a2["Lemon"] = 16
	a2["Orange"] = 44
	a2["Pineapple"] = 73
	for k, v := range a2 {
		fmt.Printf("%s => %d\n", k, v)
	}

	b1 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	fmt.Println(b1)
	fmt.Println("len:", len(b1))
	b1[4] = "D"
	b1[5] = "E"
	fmt.Println(b1)
	fmt.Println("len:", len(b1))
	delete(b1, 2)
	fmt.Println(b1)

	c1 := make(map[int]string, 100)
	fmt.Println(c1)
	fmt.Println(len(c1))
}
