package common

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var TESTDB *gorm.DB

// SetupDB : initializing mysql database
func InitDB() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")
	log.Info("DB_NAME: ", DBNAME)
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	DB = db
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

func InitTestDB() {
	USER := os.Getenv("TESTDB_USER")
	PASS := os.Getenv("TESTDB_PASS")
	HOST := os.Getenv("TESTDB_HOST")
	PORT := os.Getenv("TESTDB_PORT")
	DBNAME := os.Getenv("TESTDB_NAME")
	log.Info("TESTDB_NAME: ", DBNAME)
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	TESTDB = db
}

// Using this function to get a connection, you can create your connection pool here.
func GetTestDB() *gorm.DB {
	return TESTDB
}
