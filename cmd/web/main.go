package main

import (
	"fmt"
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/constants"
	"github.com/mjaliz/gotracktime/internal/handlers"
	"log"
	"os"
)

var app config.AppConfig
var repo *handlers.DBRepo

func init() {
	fmt.Println("Setting default timezone...")
	err := os.Setenv("TZ", constants.DefaultTimeZone)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := setupApp()
	if err != nil {
		log.Fatal(err)
	}

	r := handlers.NewRouters()
	if err = r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
