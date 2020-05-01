package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Phone struct {
	ID     int64
	Name   string
	Number string
}

func main() {
	db, err := sql.Open("sqlite3", "./phone.db")
	if err != nil {
		log.Fatal(err)
	}

	sql := `
	CREATE TABLE IF NOT EXISTS phones (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		number TEXT NOT NULL
	);
	`

	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}
}
