package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Open() {
	uri := GetSQLURI()
	db, err := sql.Open("mysql", uri)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Test connection to MySQL server
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to database.")
		return
	} else {
		fmt.Println("Successfully connected to database.")
	}

	// Expose "db" globally
	DB = db
}

func Close() error {
	return DB.Close()
}

func GetDB() *sql.DB {
	return DB
}