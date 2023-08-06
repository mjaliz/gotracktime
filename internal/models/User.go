package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Email       string `gorm:"not null;uniqueIndex"`
	Password    string `gorm:"not null"`
	UserActive  int    `gorm:"not null;default:1"`
	AccessLevel int    `gorm:"not null;default:1"`
}
