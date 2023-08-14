package dbrepo

import (
	"github.com/mjaliz/gotracktime/internal/models"
)

func (p *postgresDBRepo) InsertUser(u models.SignUpInput) (models.User, error) {
	userDB := models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	if err := p.DB.Create(&userDB).Error; err != nil {
		return models.User{}, err
	}
	return userDB, nil
}

func (p *postgresDBRepo) FindUserByEmail(u models.SignInInput) (models.User, error) {
	var user models.User
	if err := p.DB.Where("email  = ?", u.Email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
