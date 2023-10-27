package db

import (
	"database/sql"
	"echo-rest/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	// call function named getConfig from package config
	conf := config.GetConfig()

	// defining connection string using standardized formula
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	// trying to access sql using the previous config
	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("connectionString error...")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN invalid")
	}
}

func createCon() *sql.DB {
	return db
}
