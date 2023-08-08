package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/driver"
	"github.com/mjaliz/gotracktime/internal/helpers"
	"github.com/mjaliz/gotracktime/internal/inputs"
	"github.com/mjaliz/gotracktime/internal/repository"
	"github.com/mjaliz/gotracktime/internal/repository/dbrepo"
	"log"
	"net/http"
	"strings"
)

// Repo is the repository
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

func (repo *DBRepo) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (repo *DBRepo) SignUp(c *gin.Context) {
	var user inputs.User
	if err := c.ShouldBindJSON(&user); err != nil {
		validationErrs := helpers.ParseValidationError(err)
		helpers.FailedResponse(c, http.StatusBadRequest, validationErrs, "")
		return
	}
	err := repo.DB.InsertUser(user)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "duplicate key value") {
			helpers.FailedResponse(c, http.StatusBadRequest, nil, "email already exists")
		}
		return
	}
	helpers.SuccessResponse(c, http.StatusCreated, user, "")
}
