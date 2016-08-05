package repository

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDBConnection() {
	_db, err := gorm.Open("mysql", "b2e317c6245888:82f9d3be@tcp(us-cdbr-iron-east-04.cleardb.net:3306)/heroku_036792222f30804?parseTime=true")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db = _db;
}