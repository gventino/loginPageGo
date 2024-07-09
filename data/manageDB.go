package data

import (
	"database/sql"
	"fmt"
	"main/models"

	_ "github.com/mattn/go-sqlite3"
)

type User = models.User

func InsertUser(username, email, phone, password string) {
	// Open db:
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
	}

	// Insert user into db:
	statement, err := db.Prepare(`
	INSERT INTO users (username, email, phone, password)
	VALUES (?, ?, ?, ?);
	`)
	if err != nil {
		fmt.Println(err)
	}
	statement.Exec(username, email, phone, password)

	// Close db:
	defer db.Close()
}

func VerifyLogin(username, password string) bool {
	// Open db:
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
	}

	rows, _ := db.Query("SELECT username, password FROM users")
	var temp User
	for rows.Next() {
		rows.Scan(&temp.Username, &temp.Password)
		if username == temp.Username && password == temp.Password {
			return true
		}
	}

	// Close db:
	defer db.Close()

	return false
}
