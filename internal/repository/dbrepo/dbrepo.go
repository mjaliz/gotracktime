package dbrepo

import (
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/repository"
	"gorm.io/gorm"
)

var app *config.AppConfig

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *gorm.DB
}

// NewPostgresRepo creates the repository
func NewPostgresRepo(Conn *gorm.DB, a *config.AppConfig) repository.DatabaseRepo {
	app = a
	return &postgresDBRepo{
		App: a,
		DB:  Conn,
	}
}
