package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/constants"
	"github.com/mjaliz/gotracktime/internal/driver"
	"github.com/mjaliz/gotracktime/internal/helpers"
	"github.com/mjaliz/gotracktime/internal/inputs"
	"github.com/mjaliz/gotracktime/internal/repository"
	"github.com/mjaliz/gotracktime/internal/repository/dbrepo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
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
	helpers.SuccessResponse(c, http.StatusOK, nil, "Welcome to timemyth")
}

func (repo *DBRepo) SignUp(c *gin.Context) {
	var user inputs.User
	if err := c.ShouldBindJSON(&user); err != nil {
		validationErrs := helpers.ParseValidationError(err)
		helpers.FailedResponse(c, http.StatusBadRequest, validationErrs, "")
		return
	}
	if user.Password != user.PasswordConfirm {
		helpers.FailedResponse(c, http.StatusBadRequest, nil, "password and password confirm didn't match")
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
	helpers.SuccessResponse(c, http.StatusCreated, user.PrivateUser(), "")
}

func (repo *DBRepo) SignIn(c *gin.Context) {
	var user inputs.UserSignIn
	if err := c.ShouldBindJSON(&user); err != nil {
		validationErrs := helpers.ParseValidationError(err)
		helpers.FailedResponse(c, http.StatusBadRequest, validationErrs, "")
		return
	}
	userDB, err := repo.DB.FindUserByEmail(user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.FailedResponse(c, http.StatusUnauthorized, nil, "")
			return
		}
		helpers.FailedResponse(c, http.StatusInternalServerError, nil, "")
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		helpers.FailedResponse(c, http.StatusUnauthorized, nil, "")
		return
	}
	expiredAt := time.Now().UTC().Add(constants.JWTExpireDuration)
	accessToken, err := helpers.GenerateJWT(&userDB, expiredAt)
	if err != nil {
		helpers.FailedResponse(c, http.StatusInternalServerError, nil, "")
		return
	}
	c.SetCookie("accessToken", accessToken, expiredAt.Second(), "/", "localhost", false, false)
}
