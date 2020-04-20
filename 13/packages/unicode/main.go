package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(unicode.IsDigit('X'))
	fmt.Println(unicode.IsDigit('3'))
	fmt.Println(unicode.IsDigit('３'))
	fmt.Println(unicode.IsDigit('三'))

	fmt.Println(unicode.IsLetter('A'))
	fmt.Println(unicode.IsLetter('Ａ'))
	fmt.Println(unicode.IsLetter('3'))
	fmt.Println(unicode.IsLetter('３'))
	fmt.Println(unicode.IsLetter('あ'))
	fmt.Println(unicode.IsLetter('🍣'))

	fmt.Println(unicode.IsSpace(' '))
	fmt.Println(unicode.IsSpace('　'))
	fmt.Println(unicode.IsSpace('\t'))
	fmt.Println(unicode.IsSpace('_'))

	fmt.Println(unicode.IsControl('\n'))
}
