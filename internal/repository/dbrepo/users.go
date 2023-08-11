package dbrepo

import (
	"github.com/mjaliz/gotracktime/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (p *postgresDBRepo) InsertUser(u models.SignUpInput) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return models.User{}, err
	}
	userDB := models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: string(hashedPassword),
	}
	if err = p.DB.Create(&userDB).Error; err != nil {
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
