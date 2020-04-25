package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(-12345, 10))
	fmt.Println(strconv.FormatUint(12345, 10))
	fmt.Println(strconv.Itoa(3333))

	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.ParseBool("T"))
	fmt.Println(strconv.ParseBool("TRUE"))
	fmt.Println(strconv.ParseBool("True"))
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseBool("0"))
	fmt.Println(strconv.ParseBool("f"))
	fmt.Println(strconv.ParseBool("F"))
	fmt.Println(strconv.ParseBool("FALSE"))
	fmt.Println(strconv.ParseBool("False"))
	if _, err := strconv.ParseBool("Foo"); err != nil {
		log.Println(err)
	}

	fmt.Println(strconv.ParseInt("12345", 10, 0))
	fmt.Println(strconv.ParseInt("-12", 10, 0))
	fmt.Println(strconv.ParseUint("12345", 10, 0))
	if _, err := strconv.ParseUint("-100", 19, 0); err != nil {
		log.Println(err)
	}
	fmt.Println(strconv.Atoi("12345"))
}
