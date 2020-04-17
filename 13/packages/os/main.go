package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// defer func() {
	//      fmt.Println("ここは実行されない")
	// }()
	// os.Exit(0)

	_, err := os.Open("fooooooooooooooo")
	if err != nil {
		// log.Fatal(err)
	}

	fmt.Printf("length=%d\n", len(os.Args))
	for _, v := range os.Args {
		fmt.Println(v)
	}

	f1, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	bs1 := make([]byte, 128)
	n1, err := f1.Read(bs1)
	fmt.Println(n1)
	fmt.Println(string(bs1))
	bs2 := make([]byte, 128)
	n2, err := f1.ReadAt(bs2, 9)
	fmt.Println(n2)
	fmt.Println(string(bs2))
	offset1, err := f1.Seek(10, os.SEEK_SET)
	bs3 := make([]byte, 128)
	f1.Read(bs3)
	fmt.Println(offset1)
	fmt.Println(string(bs3))
	offset2, err := f1.Seek(-2, os.SEEK_CUR)
	bs4 := make([]byte, 128)
	f1.Read(bs4)
	fmt.Println(offset2)
	fmt.Println(string(bs4))
	offset3, err := f1.Seek(0, os.SEEK_END)
	bs5 := make([]byte, 128)
	f1.Read(bs5)
	fmt.Println(offset3)
	fmt.Println(string(bs5))

	fi1, err := f1.Stat()
	fmt.Println(fi1.Name())
	fmt.Println(fi1.Size())
	fmt.Println(fi1.Mode())
	fmt.Println(fi1.ModTime())
	fmt.Println(fi1.IsDir())

	f2, err := os.Create("hoge.txt")
	if err != nil {
		log.Fatal(err)
	}
	fi2, err := f2.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fi2.Name())
	fmt.Println(fi2.Size())
	fmt.Println(fi2.IsDir())
	f2.Write([]byte("Hello, World!\n"))
	f2.WriteAt([]byte("Golang"), 7)
	f2.Seek(0, os.SEEK_END)
	f2.WriteString("Yeah!")

	f3, err := os.OpenFile("foo.txt", os.O_RDONLY, 0666)
	fmt.Println(f3)

	if err := os.Remove("hoge.txt"); err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll("foo/fuga", 0775); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll("foo"); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Create("test.txt"); err != nil {
		log.Fatal(err)
	}
	if err := os.Rename("test.txt", "testdayo.txt"); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir("foo", 0755); err != nil {
		log.Fatal(err)
	}

	if err := os.Rename("foo", "bar"); err != nil {
		log.Fatal(err)
	}

	if err := os.Rename("testdayo.txt", "bar/test.txt"); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll("bar"); err != nil {
		log.Fatal(err)
	}

	here, err := os.Getwd()
	fmt.Println(here)
	if err := os.Chdir("/tmp"); err != nil {
		log.Fatal(err)
	}
	dir, err := os.Getwd()
	fmt.Println(dir)
	f4, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}
	defer f4.Close()

	fis, err := f4.Readdir(0)
	for _, fi := range fis {
		if fi.IsDir() {
			fmt.Println(fi.Name())
		}
	}

	if err := os.Chdir(here); err != nil {
		log.Fatal(err)
	}

	fmt.Println(os.TempDir())

	if err := os.Symlink("foo.txt", "bar.txt"); err != nil {
		log.Fatal(err)
	}

	path, err := os.Readlink("bar.txt")
	fmt.Println(path)

	if err := os.Remove("bar.txt"); err != nil {
		log.Fatal(err)
	}
}
