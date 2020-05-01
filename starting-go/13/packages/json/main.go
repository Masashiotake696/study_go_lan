package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id      int
	Name    string
	Email   string
	Created time.Time
}

func main() {
	u1 := new(User)
	u1.Id = 1
	u1.Name = "山田太郎"
	u1.Email = "yamada@example.com"
	u1.Created = time.Now()
	bs1, err := json.Marshal(u1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs1))

	u2 := &User{
		Id:      2,
		Name:    "鈴木次郎",
		Email:   "suzuki@example.com",
		Created: time.Now(),
	}
	bs2, err := json.Marshal(u2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs2))

	src := `
{
"Id": 12,
"Name": "田中花子",
"Email": "tanaka@example.com",
"Created": "2020-04-21T12:52:22.153653+09:00"
}
`
	u3 := new(User)
	if err := json.Unmarshal([]byte(src), u3); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", u3)
}
