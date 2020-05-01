package main

func swap(i, j int) (int, int) {
	return j, i
}

func main() {
	n, m := swap(10, 20)
	println(n, m)
}
