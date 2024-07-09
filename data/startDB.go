package data

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func StartDB() {
	// Opening db:
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
	}

	// Creating table:
	statement, err := db.Prepare(`
	CREATE TABLE users IF NOT EXISTS(
		username TEXT PRIMARY KEY,
		email TEXT UNIQUE,
		password TEXT,
		phone TEXT
	);
	`)
	if err != nil {
		fmt.Println(err)
	}
	statement.Exec()

	// Closing db:
	defer db.Close()
}
