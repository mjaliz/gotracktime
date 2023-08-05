package config

import (
	"github.com/mjaliz/gotracktime/internal/driver"
)

type AppConfig struct {
	DB *driver.DB
}
