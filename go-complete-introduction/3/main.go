package main

func main() {
	var sum int
	ns := [...]int{19, 86, 1, 12}

	for i := 0; i < len(ns); i++ {
		sum += ns[i]
	}

	println(sum)
}
