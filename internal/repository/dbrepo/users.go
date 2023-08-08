package dbrepo

import (
	"github.com/mjaliz/gotracktime/internal/inputs"
	"github.com/mjaliz/gotracktime/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (p *postgresDBRepo) InsertUser(u inputs.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
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

func (p *postgresDBRepo) FindUserByEmail(u inputs.UserSignIn) (models.User, error) {
	var user models.User
	if err := p.DB.Where("email  = ?", u.Email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
