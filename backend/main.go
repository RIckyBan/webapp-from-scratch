package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbInit() {
	dsn := "user:password@tcp(localhost:3306)/webapp?charset=utf8&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS webapp")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("CREATE DATABASE webapp") //
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	dbInit()
}
