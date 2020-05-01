package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func infoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<title>インフォメーション</title>
</head>
<body>
<h1>ようこそ！</h1>
</body>
</html>
`)
}

func main() {
	res1, err := http.Get("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res1.StatusCode)
	fmt.Println(res1.Proto)
	fmt.Println(res1.Header["Date"])
	fmt.Println(res1.Header["Content-Type"])
	fmt.Println(res1.Request.Method)
	fmt.Println(res1.Request.URL)
	defer res1.Body.Close()
	body1, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body1))

	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())

	res2, err := http.PostForm("https://example.com/comments/post", vs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res2.StatusCode)
	defer res2.Body.Close()
	body2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body2))

	f3, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	res3, err := http.Post("https://example.com/upload", "image/jpeg", f3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res3.StatusCode)
	body3, err := ioutil.ReadAll(res3.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body3))

	http.HandleFunc("/info", infoHandler)
	http.ListenAndServe(":8080", nil)
}
