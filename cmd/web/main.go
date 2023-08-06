package main

import (
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/handlers"
	"log"
)

var app config.AppConfig
var repo *handlers.DBRepo

func main() {
	err := setupApp()
	if err != nil {
		log.Fatal(err)
	}

	r := setupRoutes()
	if err = r.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
