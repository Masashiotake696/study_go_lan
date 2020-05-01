package main

// パッケージ変数
var packageVariable = 100

func one() int {
	packageVariable = packageVariable + 1
	return 1
}

func main() {
	// ローカル変数
	var n int
	var x, y, z int

	var (
		a    int
		name string
	)
	n = 5
	a = 3
	x, y, z = 3, 4, 5
	name = "Taro"
	println(n, a, x, y, z, name)

	b := true
	i := 1
	f := 3.14
	s := "abc"
	one := one()
	println(b, i, f, s, one)

	var (
		int    = 1
		string = "string"
		bool   = true
	)
	println(int, string, bool)

	packageVariable = packageVariable + 1
	println(packageVariable)
}
