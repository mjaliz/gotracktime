package models

import (
	"gorm.io/gorm"
	"time"
)

// TimeEntity is the model for each time records
type TimeEntity struct {
	gorm.Model
	StartedAt     time.Time `gorm:"not null"`
	StoppedAt     *time.Time
	UserID        uint `gorm:"not null"`
	DescriptionID *uint
	ProjectID     *uint
}

type TimeEntityInput struct {
	UserID        uint      `json:"-"`
	StartedAt     time.Time `json:"started_at" binding:"required"`
	DescriptionID *uint     `json:"description_id,omitempty"`
	ProjectID     *uint     `json:"project_id,omitempty"`
}

type TimeEntityOutput struct {
	TimeEntityInput
	CreatedAt time.Time `json:"created_at"`
}
