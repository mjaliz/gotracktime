package driver

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	SQL *gorm.DB
}

var dbConn = &DB{}

// ConnectPostgres create database pool for postgres
func ConnectPostgres(dsn string) (*DB, error) {
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
}
