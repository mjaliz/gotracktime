package models

import "gorm.io/gorm"

// Description is the model for TimeEntity description
type Description struct {
	gorm.Model
	Text         string
	TimeEntities []TimeEntity
}
