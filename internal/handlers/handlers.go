package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/driver"
	"github.com/mjaliz/gotracktime/internal/repository"
	"github.com/mjaliz/gotracktime/internal/repository/dbrepo"
	"net/http"
)

//Repo is the repository
var Repo *DBRepo
var app *config.AppConfig

// DBRepo is the db repo
type DBRepo struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewHandlers creates the handlers
func NewHandlers(repo *DBRepo, a *config.AppConfig) {
	Repo = repo
	app = a
}

// NewPostgresqlHandlers creates db repo for postgres
func NewPostgresqlHandlers(db *driver.DB, a *config.AppConfig) *DBRepo {
	return &DBRepo{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func (repo *DBRepo) Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
