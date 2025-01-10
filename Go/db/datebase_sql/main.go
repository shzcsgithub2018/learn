package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 打开或创建 SQLite 数据库
	db, err := sql.Open("sqlite3", "your_database.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// 执行一些数据库操作，例如创建表
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INT PRIMARY KEY,
        name VARCHAR(255)
    );`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("SQLite database initialized successfully!")

	var (
		id   int
		name string
	)

	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
