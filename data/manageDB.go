package data

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InsertUser(username, email, phone, password string) bool {
	// Open db:
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT username FROM users WHERE username=?", username)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// If user not registered yet
	if !rows.Next() {
		// Insert user into db:
		statement, err := db.Prepare(`
		INSERT INTO users (username, email, phone, password)
		VALUES (?, ?, ?, ?);
		`)
		if err != nil {
			fmt.Println(err)
		}
		statement.Exec(username, email, phone, password)

		defer statement.Close()
		return true
	}

	return false
}

func VerifyLogin(username, password string) bool {
	// Open db:
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT username, password FROM users WHERE username=? AND password=?", username, password)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	return rows.Next()
}

func ResetDB() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM users")
	if err != nil {
		fmt.Println(err)
	}
	statement.Exec()
	defer statement.Close()
}
