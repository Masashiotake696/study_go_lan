package main

import (
	"fmt"
)

func main() {
	// deferはpanicを使っても実行される
	defer fmt.Println("on defer")
	// panicは処理を終了させる。だたし、deferは実行される
	panic("runtime error!")
	// os.Exit(0)は処理を終了させる。ただし、deferは実行されない
	// os.Exit(0)
	fmt.Println("Hello, world")
}
