package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

var DB *sql.DB

func InitDB() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	var err error

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("error openning database: ", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	//defer db.Close()

	fmt.Println("successfully connected to mariadb...")

	createTables()

}

func createTables() {
	createPostTable := `
		CREATE TABLE IF NOT EXISTS posts(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		title TEXT NOT NULL,
		text TEXT NOT NULL,
		time DATETIME NOT NULL
		)
	`
	_, err := DB.Exec(createPostTable)

	if err != nil {
		panic("cant create posts table")
	}

	createUserTable := `
		CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name TEXT NOT NULL,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
		)
		`

	_, err = DB.Exec(createUserTable)

	if err != nil {
		panic("cant create users table")
	}
}
