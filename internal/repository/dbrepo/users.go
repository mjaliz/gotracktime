package dbrepo

import (
	"github.com/mjaliz/gotracktime/internal/models"
	"github.com/mjaliz/gotracktime/internal/utils"
	"gorm.io/gorm"
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

func (p *testDBRepo) InsertUser(u models.SignUpInput) (models.User, error) {
	userDB := models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	return userDB, nil
}

func (p *testDBRepo) FindUserByEmail(u models.SignInInput) (models.User, error) {
	hashedPassword, err := utils.HashPassword(u.Password)
	var user models.User
	user.Email = u.Email
	user.Password = hashedPassword
	if err != nil {
		return models.User{}, err
	}
	if u.Email == "unauthorized@gmail.com" {
		return models.User{}, gorm.ErrRecordNotFound
	}
	if u.Email == "internalServerError@gmail.com" {
		return models.User{}, gorm.ErrUnsupportedDriver
	}
	if u.Email == "hashPasswordError@gmail.com" {
		user.Password = "wrongPassword"
		return user, nil
	}

	return user, nil
}
