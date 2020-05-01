package main

import (
	"fmt"

	"github.com/jinzhu/configor"
)

var Config = struct {
	DB string `env:"DB"`
}{}

func main() {
	configor.Load(&Config)
	fmt.Printf("config: %#v", Config)
}
