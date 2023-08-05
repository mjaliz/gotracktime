package helpers

import "github.com/mjaliz/gotracktime/internal/config"

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}
