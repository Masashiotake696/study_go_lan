package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join([]string{"A", "B", "C"}, "/"))
	fmt.Println(strings.Join([]string{"Hello", ", ", "World!"}, ""))
	fmt.Println(strings.Join([]string{"Hello", "World!"}, ", "))

	t1 := "ABCDE"
	fmt.Println(strings.Index(t1, "A"))
	fmt.Println(strings.Index(t1, "BCD"))
	fmt.Println(strings.Index(t1, "DEF"))
	t2 := "ABCABCABC"
	fmt.Println(strings.Index(t2, "ABC"))
	fmt.Println(strings.LastIndex(t2, "ABC"))

	fmt.Println(strings.IndexAny("ABC", "ABC"))
	fmt.Println(strings.IndexAny("ABC", "BCD"))
	fmt.Println(strings.IndexAny("ABC", "CDE"))
	fmt.Println(strings.IndexAny("ABC", "XYZ"))
	fmt.Println(strings.LastIndexAny("ABC", "ABC"))
	fmt.Println(strings.LastIndexAny("ABC", "XYZ"))

	fmt.Println(strings.HasPrefix("Go言語", "Go"))
	fmt.Println(strings.HasPrefix("Go言語", "言語"))
	fmt.Println(strings.HasSuffix("Go言語", "Go"))
	fmt.Println(strings.HasSuffix("Go言語", "言語"))

	fmt.Println(strings.Contains("ABCDE", "AB"))
	fmt.Println(strings.Contains("ABCDE", "CDE"))
	fmt.Println(strings.Contains("ABCDE", "XYZ"))
	fmt.Println(strings.Contains("ABCDE", ""))
	fmt.Println(strings.Contains("", ""))

	fmt.Println(strings.ContainsAny("ABCDE", "AE"))
	fmt.Println(strings.ContainsAny("ABCDE", "Cookbook"))
	fmt.Println(strings.ContainsAny("ABCDE", "XYZ"))

	fmt.Println(strings.Count("ABC", "ABC"))
	fmt.Println(strings.Count("ABCABCABC", "ABC"))
	fmt.Println(strings.Count("ABCABCABC", "XYZ"))
	fmt.Println(strings.Count("ABC", ""))

	fmt.Println(strings.Repeat("ABC", 3))
	fmt.Println(strings.Repeat("ABC", 0))

	fmt.Println(strings.Replace("AAAAA", "A", "X", 1))
	fmt.Println(strings.Replace("AAAAA", "A", "X", 2))
	fmt.Println(strings.Replace("AAAAA", "A", "X", -1))
	fmt.Println(strings.Replace("C言語", "C", "Go", 1))

	fmt.Printf("%#v\n", strings.Split("A,B,C,D,E", ","))
	fmt.Printf("%#v\n", strings.SplitAfter("A,B,C,D,E", ","))
	fmt.Printf("%#v\n", strings.SplitN("A,B,C,D,E", ",", 3))
	fmt.Printf("%#v\n", strings.SplitAfterN("A,B,C,D,E", ",", 3))

	fmt.Println(strings.ToLower("ABCDE"))
	fmt.Println(strings.ToLower("X Y Z"))
	fmt.Println(strings.ToUpper("abcde"))
	fmt.Println(strings.ToUpper("x y z"))

	fmt.Println(strings.TrimSpace(" - Hello, World! - "))
	fmt.Println(strings.TrimSpace("\tGolang\r\n"))

	fmt.Printf("%#v\n", strings.Split("A B C D E", " "))
	fmt.Printf("%#v\n", strings.Fields("A B C D E"))
}
