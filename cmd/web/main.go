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

	e := routes()
	e.Logger.Fatal(e.Start(":1323"))
}
