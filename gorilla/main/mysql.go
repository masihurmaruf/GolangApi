package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@unix(/Applications/MAMP/tmp/mysql/mysql.sock)/go-db")
	if err != nil {
		log.Fatal("Can not connect!")
	}
	return db
}
