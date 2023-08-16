package models

import (
	"gorm.io/gorm"
	"time"
)

// User is the user model definition
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;uniqueIndex"`
	Password     string `gorm:"not null"`
	Role         int    `gorm:"not null;default:1"`
	IsActive     bool   `gorm:"not null;default:false"`
	TimeEntities []TimeEntity
	Projects     []Project
}

type SignUpInput struct {
	Name            string `json:"name,omitempty"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpOutput struct {
	UID       uint      `json:"uid"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type SignInOutput struct {
	AccessToken string `json:"access_token"`
}

func (u *User) FilterUserResponse() SignUpOutput {
	return SignUpOutput{
		UID:       u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
