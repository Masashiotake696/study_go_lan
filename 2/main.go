package main

import (
	"./animal"
	"fmt"
)

func main() {
	// AppNameを実行
	fmt.Println(AppName());
	// animalパッケージの関数を実行
	fmt.Println(animal.ElephantFeed());
	fmt.Println(animal.MonkeyFeed());
	fmt.Println(animal.RabbitFeed());
}