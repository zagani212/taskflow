package main

import (
	"database/sql"
  	"fmt"
  	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "taskflow"
)

var db *sql.DB

func CreateConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	conn, err := sql.Open("postgres", psqlInfo)
	db = conn
	if err != nil{
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Here")
		panic(err)
	}
}

func CloseConnection() {
	db.Close()
}