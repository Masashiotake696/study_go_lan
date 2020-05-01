package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	convertimage "./convert_image"
)

var (
	b string
	a string
)

func init() {
	flag.StringVar(&b, "b", convertimage.ExtJpeg, "変換前画像形式")
	flag.StringVar(&a, "a", convertimage.ExtPng, "変換後画像形式")
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("対象ディレクトリを指定してください")
		os.Exit(1)
	}

	for _, dir := range flag.Args() {
		cis, err := convertimage.NewConvertImagesByDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		err = cis.ConvertImagesFromTo(b, a)
		if err != nil {
			log.Fatal(err)
		}
	}
}
