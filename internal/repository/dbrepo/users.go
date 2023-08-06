package dbrepo

import (
	"github.com/mjaliz/gotracktime/internal/inputs"
	"github.com/mjaliz/gotracktime/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (p *postgresDBRepo) InsertUser(u inputs.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}
	userDB := models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: string(hashedPassword),
	}
	if err = p.DB.Create(&userDB).Error; err != nil {
		return err
	}
	return nil
}
