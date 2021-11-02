package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	uri := get_mysql_uri()
	db, err := sql.Open("mysql", uri)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to database.")
		return
	} else {
		fmt.Println("Successfully connected to database.")
	}
}