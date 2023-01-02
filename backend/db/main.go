package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func createDSN() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true",
		os.Getenv("DBUser"),
		os.Getenv("DBPassword"),
		os.Getenv("DBHost"),
		os.Getenv("DBPort"),
		os.Getenv("DBName"),
	)

	return dsn, nil
}

func OpenDB() (*sql.DB, error) {
	dsn, err := createDSN()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDB() error {
	db, err := OpenDB()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS webapp")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE DATABASE webapp")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := initDB()
	if err != nil {
		panic(err.Error())
	}
}
