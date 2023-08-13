package dbrepo

import (
	"github.com/mjaliz/gotracktime/internal/models"
)

func (p *testDBRepo) InsertUser(u models.SignUpInput) (models.User, error) {
	userDB := models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	return userDB, nil
}

func (p *testDBRepo) FindUserByEmail(u models.SignInInput) (models.User, error) {
	var user models.User
	return user, nil
}
