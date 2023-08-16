package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/constants"
	"github.com/mjaliz/gotracktime/internal/driver"
	"github.com/mjaliz/gotracktime/internal/repository"
	"github.com/mjaliz/gotracktime/internal/repository/dbrepo"
	"github.com/mjaliz/gotracktime/internal/utils"
	"net/http"
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

func NewTestHandlers(a *config.AppConfig) *DBRepo {
	return &DBRepo{
		App: a,
	}
}

func (repo *DBRepo) Home(c *gin.Context) {
	utils.SuccessResponse(c, http.StatusOK, nil, "Welcome to timemyth")
}

func Ping(c *gin.Context) {
	user := getUserFromContext(c)
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "pong", "userid": user.UserID}, "")
}

func getUserFromContext(c *gin.Context) (user *utils.JWTClaim) {
	userClaim, ok := c.Get(constants.UserClaims)
	if !ok {
		utils.FailedResponse(c, http.StatusForbidden, nil, "")
	}
	user = userClaim.(*utils.JWTClaim)
	return
}
