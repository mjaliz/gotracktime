package models

import "gorm.io/gorm"

// Project is the model for TimeEntity project
type Project struct {
	gorm.Model
	Title        string `gorm:"not null"`
	TimeEntities []TimeEntity
}
