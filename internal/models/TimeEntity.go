package models

import (
	"gorm.io/gorm"
	"time"
)

// TimeEntity is the model for each time records
type TimeEntity struct {
	gorm.Model
	StartedAt     time.Time `gorm:"not null"`
	StoppedAt     time.Time `gorm:"not null"`
	UserID        uint      `gorm:"not null"`
	DescriptionID uint
	ProjectID     uint
}
