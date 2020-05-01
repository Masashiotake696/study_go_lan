package main

import (
	"log"
	"os"
)

func main() {
	log.Print("ログの1行目\n")
	log.Println("ログの2行目")
	log.Printf("ログの%d行目\n", 3)

	f, err := os.Create("test.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	log.Println("ログメッセージ")

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags)
	log.Println("A")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")
	log.SetFlags(log.Ltime | log.Llongfile)
	log.Println("D")
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[LOG] ")
	log.Println("E")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Fatalln("message")
}
