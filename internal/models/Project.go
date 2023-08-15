package models

import (
	"gorm.io/gorm"
	"time"
)

// Project is the model for TimeEntity project
type Project struct {
	gorm.Model
	Title        string `gorm:"not null"`
	TimeEntities []TimeEntity
}

type ProjectInput struct {
	Title string `json:"title" binding:"required"`
}

type ProjectOutput struct {
	ProjectInput
	CreatedAt time.Time `json:"created_at"`
}
