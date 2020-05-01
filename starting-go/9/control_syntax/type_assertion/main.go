package main

import "fmt"

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything(1)
	anything("a")
	anything([...]int{1, 2, 3})

	var z interface{} = 3
	fmt.Println(z, z.(int)+1)
	_, isInt := z.(int)
	fmt.Println(isInt)

	z = "test string"
	if z == nil {
		fmt.Println("x is nil")
	} else if i, isInt := z.(int); isInt {
		fmt.Printf("z is integer : %d\n", i)
	} else if s, isString := z.(string); isString {
		fmt.Printf("z is string : %s\n", s)
	} else {
		fmt.Println("unsupported type")
	}

	switch v := z.(type) {
	case bool:
		fmt.Println("bool:", v)
	case int, uint:
		fmt.Println("integer or unsigned integer:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("don't know:", v)
	}
}
