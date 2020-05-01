package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	sql := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);
	`
	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}

	users := []*User{
		{
			Name: "Otake",
			Age:  24,
		},
		{
			Name: "Goper",
			Age:  10,
		},
	}
	for _, user := range users {
		r, err := db.Exec("INSERT INTO users (name, age) values (?, ?)", user.Name, user.Age)
		if err != nil {
			log.Fatal(err)
		}
		id, err := r.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		user.ID = id
		fmt.Println("INSERT", user)
	}

	rows, err := db.Query("SElECT * FROM users WHERE age = 10")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		u := User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	r, err := db.Exec("UPDATE users SET age = age + 1 WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}
	cnt, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Affected rows: ", cnt)
}
