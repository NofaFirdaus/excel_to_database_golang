package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Database() {
	dsn := "root:nofafirdausananta@tcp(127.0.0.1:3307)/excel_to_databases"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	log.Println("databases connect")
	DB = db
}
