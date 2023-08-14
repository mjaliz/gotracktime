package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/constants"
	"github.com/mjaliz/gotracktime/internal/driver"
	"github.com/mjaliz/gotracktime/internal/models"
	"github.com/mjaliz/gotracktime/internal/repository"
	"github.com/mjaliz/gotracktime/internal/repository/dbrepo"
	"github.com/mjaliz/gotracktime/internal/utils"
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
	utils.SuccessResponse(c, http.StatusOK, nil, "Welcome to timemyth")
}

func (repo *DBRepo) SignUp(c *gin.Context) {
	var userInput models.SignUpInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		validationErrs := utils.ParseValidationError(err)
		utils.FailedResponse(c, http.StatusBadRequest, validationErrs, "")
		return
	}
	if userInput.Password != userInput.PasswordConfirm {
		utils.FailedResponse(c, http.StatusBadRequest, nil, "password and password confirm didn't match")
		return
	}
	hashedPassword, err := utils.HashPassword(userInput.Password)
	if err != nil {
		utils.FailedResponse(c, http.StatusInternalServerError, nil, "")
		return
	}
	userInput.Password = hashedPassword
	userDB, err := repo.DB.InsertUser(userInput)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "duplicate key value") {
			utils.FailedResponse(c, http.StatusBadRequest, nil, "email already exists")
		}
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, userDB.FilterUserResponse(), "")
}

func (repo *DBRepo) SignIn(c *gin.Context) {
	var userInput models.SignInInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		validationErrs := utils.ParseValidationError(err)
		utils.FailedResponse(c, http.StatusBadRequest, validationErrs, "")
		return
	}
	userDB, err := repo.DB.FindUserByEmail(userInput)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.FailedResponse(c, http.StatusUnauthorized, nil, "")
			return
		}
		utils.FailedResponse(c, http.StatusInternalServerError, nil, "")
		return
	}
	if err = utils.ComparePassword(userDB.Password, userInput.Password); err != nil {
		utils.FailedResponse(c, http.StatusUnauthorized, nil, "")
		return
	}
	expiredAt := time.Now().UTC().Add(constants.JWTExpireDuration)
	accessToken, err := utils.GenerateJWT(&userDB, expiredAt)
	if err != nil {
		utils.FailedResponse(c, http.StatusInternalServerError, nil, "")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, models.SignInOutput{AccessToken: accessToken}, "")
}

func Ping(c *gin.Context) {
	userClaim, ok := c.Get(constants.UserClaims)
	if !ok {
		utils.FailedResponse(c, http.StatusForbidden, nil, "")
	}
	user := userClaim.(*utils.JWTClaim)

	//fmt.Println()
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "pong", "userid": user.UserID}, "")
}

func (repo *DBRepo) CreateTimeEntity(c *gin.Context) {
	userClaim, ok := c.Get(constants.UserClaims)
	if !ok {
		utils.FailedResponse(c, http.StatusForbidden, nil, "")
	}
	user := userClaim.(*utils.JWTClaim)
	var timeEntityInput models.TimeEntityInput
	if err := c.ShouldBindJSON(&timeEntityInput); err != nil {
		validationErrs := utils.ParseValidationError(err)
		utils.FailedResponse(c, http.StatusBadRequest, validationErrs, "")
		return
	}
	timeEntityInput.UserID = user.UserID
	timeEntityDB, err := repo.DB.InsertTimeEntity(timeEntityInput)
	if err != nil {
		utils.FailedResponse(c, http.StatusInternalServerError, nil, "")
		return
	}
	var output models.TimeEntityOutput
	output.CreatedAt = timeEntityDB.CreatedAt
	output.StartedAt = timeEntityDB.StartedAt
	output.DescriptionID = timeEntityDB.DescriptionID
	output.ProjectID = timeEntityDB.ProjectID
	utils.SuccessResponse(c, http.StatusCreated, output, "")
}
