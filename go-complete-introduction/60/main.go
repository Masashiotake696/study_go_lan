package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := sql.Open("sqlite3", "./money_transfer.db")
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err := initial(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := send(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func initial(tx *sql.Tx) error {
	const ctbSQL = `
	CREATE TABLE IF NOT EXISTS money (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		amount INTEGER NOT NULL
	);
	`

	const istSQL = `
	INSERT INTO money (name, amount) VALUES ("user1", 100);
	INSERT INTO money (name, amount) VALUES ("user2", 20);
	`

	if _, err := tx.Exec(ctbSQL); err != nil {
		return err
	}

	if _, err := tx.Exec(istSQL); err != nil {
		return err
	}

	return nil
}

func send(tx *sql.Tx) error {
	const deSQL = `UPDATE money SET amount = amount - 10 WHERE id = 1;`
	const inSQL = `UPDATE money SET amount = amount + 10 WHERE id = 2;`

	if _, err := tx.Exec(deSQL); err != nil {
		return err
	}

	if _, err := tx.Exec(inSQL); err != nil {
		return err
	}

	return nil
}
