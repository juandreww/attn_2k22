package model

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "ec2-52-73-155-171.compute-1.amazonaws.com"
	port = 5432
	user = "'awtpbzyctlpydm'"
	password = "a7bf40c39496f73a03e7412befbc787d29138445d7fce2a34bf31df40cf07d96"
	dbname = "d4ehughfapgq0k"
)

var con *sql.DB

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=require",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("Error when connecting to DB")
		log.Fatal(err)
	}

	con = db
	return db
}