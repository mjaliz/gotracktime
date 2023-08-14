package driver

import (
	"github.com/mjaliz/gotracktime/internal/models"
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
	err = autoMigrate(d)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
}

func autoMigrate(d *gorm.DB) error {
	err := d.AutoMigrate(&models.User{}, &models.Project{}, &models.Description{}, &models.TimeEntity{})
	if err != nil {
		return err
	}
	return nil
}
