package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Phone struct {
	ID     int64
	Name   string
	Number string
}

func main() {
	fmt.Println("[電話帳アプリ]\n")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := sql.Open("sqlite3", "./phone.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := createTable(db); err != nil {
		return err
	}

	var a int
	const message = `処理を行いたい番号を入力してください
	1 : 登録
	2 : 更新
	3 : 終了
>`
	for {
		if err := showRecords(db); err != nil {
			return err
		}

		fmt.Print(message)
		switch fmt.Scan(&a); a {
		case 1:
			if err := inputRecord(db); err != nil {
				return err
			}
		case 2:
			if err := updateRecord(db); err != nil {
				return err
			}
		case 3:
			return nil
		default:
			fmt.Println("該当する番号がありません\n")
		}

	}
}

func createTable(db *sql.DB) error {
	const sql = `
	CREATE TABLE IF NOT EXISTS phones (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		number TEXT NOT NULL
	);
	`

	if _, err := db.Exec(sql); err != nil {
		return err
	}

	return nil
}

func showRecords(db *sql.DB) error {
	fmt.Println("全件表示")

	const sql = "SELECT * FROM phones"
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}

	for rows.Next() {
		p := Phone{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Number); err != nil {
			return err
		}
		fmt.Printf("[%d] Name:%s TEL:%s\n", p.ID, p.Name, p.Number)
	}
	fmt.Println("-------------------------\n")

	return nil
}

func inputRecord(db *sql.DB) error {
	fmt.Println("新規登録")

	p := Phone{}

	fmt.Println("名前を入力してください")
	fmt.Print("Name >")
	fmt.Scan(&p.Name)

	fmt.Println("電話番号を入力してください")
	fmt.Print("TEL >")
	fmt.Scan(&p.Number)

	const sql = "INSERT INTO phones (name, number) VALUES (?,?)"
	_, err := db.Exec(sql, p.Name, p.Number)
	if err != nil {
		return err
	}
	return nil
}

func updateRecord(db *sql.DB) error {
	fmt.Println("登録更新")

	p := Phone{}

	fmt.Println("更新したいレコードのIDを入力してください")
	fmt.Print("ID >")
	fmt.Scan(&p.ID)

	fmt.Println("新しい名前を入力してください")
	fmt.Print("Name >")
	fmt.Scan(&p.Name)

	fmt.Println("新しい電話番号を入力してください")
	fmt.Print("TEL >")
	fmt.Scan(&p.Number)

	const sql = "UPDATE phones SET name = ?, number = ? WHERE id = ?"
	r, err := db.Exec(sql, p.Name, p.Number, p.ID)
	if err != nil {
		return err
	}
	cnt, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		fmt.Println("\n該当するIDのレコードが存在しませんでした\n")
	}

	return nil
}
