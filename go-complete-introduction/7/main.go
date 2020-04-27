package main

func swap2(i, j *int) {
	tmp := *i
	*i = *j
	*j = tmp
}

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}
