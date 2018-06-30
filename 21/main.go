package main

// init関数はmain関数よりも先に実行される
// init関数は複数定義できる
func init() {
	println("init1")
}
func init() {
	println("init2")
}
func init() {
	println("init3")
}
func init() {
	println("init4")
}

func main() {
	println("main")
}
