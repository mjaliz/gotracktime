package repository

import (
	"github.com/mjaliz/gotracktime/internal/inputs"
	"github.com/mjaliz/gotracktime/internal/models"
)

type DatabaseRepo interface {
	InsertUser(user inputs.User) error
	FindUserByEmail(user inputs.UserSignIn) (models.User, error)
}
