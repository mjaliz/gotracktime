package utils

import (
	"github.com/mjaliz/gotracktime/internal/config"
	"github.com/mjaliz/gotracktime/internal/constants"
	"golang.org/x/crypto/bcrypt"
)

var app *config.AppConfig

func NewUtils(a *config.AppConfig) {
	app = a
}

func HashPassword(pas string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pas), constants.PasswordHashCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func ComparePassword(dbPass, inputPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(inputPass))
}
