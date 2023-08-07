package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/driver"
	"github.com/mjaliz/gotracktime/internal/inputs"
	"github.com/mjaliz/gotracktime/internal/repository"
	"github.com/mjaliz/gotracktime/internal/repository/dbrepo"
	"log"
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

func (repo *DBRepo) Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	var errMsg string
	switch fe.Tag() {
	case "required":
		errMsg = "This field is required"
	case "lte":
		errMsg = fmt.Sprintf("Should be less than %s", fe.Param())
	case "gte":
		errMsg = fmt.Sprintf("Should be greater than %s", fe.Param())
	default:
		errMsg = "Unknown error"
	}
	return errMsg
}

func (repo *DBRepo) SignUp(c *gin.Context) {
	var user inputs.User
	if err := c.ShouldBindJSON(&user); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": out,
				"data":    nil,
			})
		}
		return
	}
	err := repo.DB.InsertUser(user)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": nil,
		"data":    user,
	})
}
