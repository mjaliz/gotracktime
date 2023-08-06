package main

import (
	"fmt"
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/driver"
	"github.com/mjaliz/gotracktime/internal/handlers"
	"github.com/mjaliz/gotracktime/internal/helpers"
	"log"
	"os"
)

func setupApp() error {
	log.Println("Connecting to database....")

	dsnString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	db, err := driver.ConnectPostgres(dsnString)
	if err != nil {
		log.Fatal("Cannot connect to database!", err)
	}

	a := config.AppConfig{
		DB: db,
	}

	app = a

	repo = handlers.NewPostgresqlHandlers(db, &app)
	handlers.NewHandlers(repo, &app)
	helpers.NewHelpers(&app)
	return err
}
