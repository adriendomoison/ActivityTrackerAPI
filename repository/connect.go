package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitDBConnection() {
	var err error
	db, err = sql.Open("mysql", "b2e317c6245888:82f9d3be@tcp(us-cdbr-iron-east-04.cleardb.net:3306)/heroku_036792222f30804")
	if err != nil {
		log.Println(err)
	}
	if err = db.Ping(); err != nil {
		log.Println(err)
	}
}